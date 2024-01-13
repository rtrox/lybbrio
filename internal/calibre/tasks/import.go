package tasks

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"lybbrio/internal/calibre"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/schema/filetype"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/scheduler"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
)

type importOutputCtxKey string

const importOutputKey importOutputCtxKey = "importOutput"

var ignorableFilenames []string = []string{
	"cover",
	"metadata",
}

type importContext struct {
	visitedAuthors    map[int64]ksuid.ID
	visitedTags       map[int64]ksuid.ID
	visitedPublishers map[int64]ksuid.ID
	visitedLanguages  map[int64]ksuid.ID
	visitedSeries     map[int64]ksuid.ID
	failedBooks       []string
}

func newImportContext() *importContext {
	return &importContext{
		visitedAuthors:    make(map[int64]ksuid.ID),
		visitedTags:       make(map[int64]ksuid.ID),
		visitedPublishers: make(map[int64]ksuid.ID),
		visitedLanguages:  make(map[int64]ksuid.ID),
		visitedSeries:     make(map[int64]ksuid.ID),
	}
}

func (c *importContext) String() string {
	var ret strings.Builder
	if len(c.failedBooks) > 0 {
		ret.WriteString("Failed Books:\n")
		for _, book := range c.failedBooks {
			ret.WriteString(fmt.Sprintf("\t%s\n", book))
		}
	}
	return ret.String()
}

func (c *importContext) AddFailedBook(book string) {
	c.failedBooks = append(c.failedBooks, book)
}

func (c *importContext) AuthorVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedAuthors[id]
	return ret, ok
}

func (c *importContext) AddAuthorVisited(id int64, ksuid ksuid.ID) {
	c.visitedAuthors[id] = ksuid
}

func (c *importContext) TagVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedTags[id]
	return ret, ok
}

func (c *importContext) AddTagVisited(id int64, ksuid ksuid.ID) {
	c.visitedTags[id] = ksuid
}

func (c *importContext) PublisherVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedPublishers[id]
	return ret, ok
}

func (c *importContext) AddPublisherVisited(id int64, ksuid ksuid.ID) {
	c.visitedPublishers[id] = ksuid
}

func (c *importContext) LanguageVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedLanguages[id]
	return ret, ok
}

func (c *importContext) AddLanguageVisited(id int64, ksuid ksuid.ID) {
	c.visitedLanguages[id] = ksuid
}

func (c *importContext) SeriesVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedSeries[id]
	return ret, ok
}

func (c *importContext) AddSeriesVisited(id int64, ksuid ksuid.ID) {
	c.visitedSeries[id] = ksuid
}

func importContextFrom(ctx context.Context) *importContext {
	output := ctx.Value(importOutputKey)
	if output == nil {
		return nil
	}
	return output.(*importContext)
}

func importContextTo(ctx context.Context, output *importContext) context.Context {
	return context.WithValue(ctx, importOutputKey, output)
}

func ImportTask(cal calibre.Calibre, client *ent.Client) scheduler.TaskFunc {
	return func(ctx context.Context, task *ent.Task, cb scheduler.ProgressCallback) (string, error) {
		log := log.Ctx(ctx)
		log.Info().Interface("task", task.ID.String()).Msg("ImportTask")

		ic := newImportContext()
		ctx = importContextTo(ctx, ic)

		err := ImportBooks(cal, client, ctx, cb)
		if err != nil {
			return "", err
		}

		return ic.String(), nil
	}
}

func ImportBooks(cal calibre.Calibre, client *ent.Client, ctx context.Context, cb scheduler.ProgressCallback) error {
	books, err := cal.GetBooks(ctx)
	if err != nil {
		return err
	}
	total := len(books)
	for idx, book := range books {
		ic := importContextFrom(ctx)

		bookCreate := client.Book.Create().
			SetTitle(book.Title).
			SetSort(book.Sort).
			SetCalibreID(book.ID).
			SetIsbn(book.ISBN).
			SetPath(book.Path).
			SetDescription(book.Comments.Text)
		if book.PubDate != nil {
			bookCreate.SetPublishedDate(*book.PubDate)
		}
		if book.SeriesIndex != nil {
			bookCreate.SetSeriesIndex(*book.SeriesIndex)
		}

		entBook, err := bookCreate.
			Save(ctx)
		if err != nil {
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Str("book", book.Title).
					Int64("bookID", book.ID).
					Msg("Book already exists")
			} else {
				log.Warn().Err(err).
					Str("book", book.Title).
					Int64("bookID", book.ID).
					Msg("Failed to create book")
				ic.AddFailedBook(book.Title)
			}
			if err := cb(float64(idx+1) / (float64(total))); err != nil {
				log.Warn().Err(err).
					Str("book", book.Title).
					Int64("bookID", book.ID).
					Msg("Failed to update progress")
			}
			continue
		}

		err = createOrAttachAuthors(client, ctx, entBook, book.Authors)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create authors")
		}

		err = createIdentifiers(ctx, client, entBook, book.Identifiers)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create identifiers")
		}

		err = createOrAttachTags(ctx, client, entBook, book.Tags)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create tags")
		}

		err = createOrAttachPublishers(ctx, client, entBook, book.Publisher)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create publishers")
		}

		err = createOrAttachLanguages(ctx, client, entBook, book.Languages)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create languages")
		}

		err = createOrAttachSeriesList(ctx, client, entBook, book.Series)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create series")
		}

		err = registerBookFiles(ctx, cal, client, entBook, *book)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed register book files")
			ic.AddFailedBook(book.Title)
		}

		if err := cb(float64(idx+1) / (float64(total))); err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to update progress")
		}
	}

	return nil
}

func createOrAttachAuthors(client *ent.Client, ctx context.Context, book *ent.Book, authors []calibre.Author) error {
	log := log.Ctx(ctx).With().Str("book", book.Title).Str("bookID", book.ID.String()).Logger()
	for _, a := range authors {
		err := createOrAttachAuthor(ctx, client, book, a)
		if err != nil {
			log.Warn().Err(err).
				Str("author", a.Name).
				Msg("Failed to create/attach author.")
		}
	}
	return nil
}

func createOrAttachAuthor(ctx context.Context, client *ent.Client, book *ent.Book, author calibre.Author) error {
	ic := importContextFrom(ctx)
	if ksuid, ok := ic.AuthorVisited(author.ID); ok {
		if err := book.Update().
			AddAuthorIDs(ksuid).
			Exec(ctx); err != nil {
			return fmt.Errorf("Error adding author to book: %w", err)
		}
		return nil
	}
	newAuthor, err := client.Author.Create().
		SetName(author.Name).
		SetSort(author.Sort).
		SetLink(author.Link).
		SetCalibreID(author.ID).
		AddBooks(book).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create author (%d): %w", author.ID, err)
	}
	ic.AddAuthorVisited(author.ID, newAuthor.ID)
	return nil
}

func createIdentifiers(ctx context.Context, client *ent.Client, book *ent.Book, identifiers []calibre.Identifier) error {

	for _, i := range identifiers {
		err := createIdentifier(ctx, client, book, i)
		if err != nil {
			log.Warn().Err(err).
				Str("identifier", i.Val).
				Msg("Failed to create identifier")
		}
	}
	return nil
}

func createIdentifier(ctx context.Context, client *ent.Client, book *ent.Book, identifier calibre.Identifier) error {
	err := client.Identifier.Create().
		SetType(identifier.Type).
		SetValue(identifier.Val).
		SetBook(book).
		SetCalibreID(identifier.ID).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func createOrAttachTags(ctx context.Context, client *ent.Client, book *ent.Book, tags []calibre.Tag) error {
	for _, t := range tags {
		err := createOrAttachTag(ctx, client, book, t)
		if err != nil {
			log.Warn().Err(err).
				Str("tag", t.Name).
				Msg("Failed to create/attach tag.")
		}
	}
	return nil
}

func createOrAttachTag(ctx context.Context, client *ent.Client, book *ent.Book, tag calibre.Tag) error {
	ic := importContextFrom(ctx)
	if ksuid, ok := ic.TagVisited(tag.ID); ok {
		if err := book.Update().
			AddTagIDs(ksuid).
			Exec(ctx); err != nil {
			return fmt.Errorf("Error adding tag to book: %w", err)
		}
		return nil
	}
	newTag, err := client.Tag.Create().
		SetName(tag.Name).
		SetCalibreID(tag.ID).
		AddBooks(book).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create tag (%d): %w", tag.ID, err)
	}
	ic.AddTagVisited(tag.ID, newTag.ID)
	return nil
}

func createOrAttachPublishers(ctx context.Context, client *ent.Client, book *ent.Book, publishers []calibre.Publisher) error {
	for _, p := range publishers {
		err := createOrAttachPublisher(ctx, client, book, p)
		if err != nil {
			log.Warn().Err(err).
				Str("publisher", p.Name).
				Msg("Failed to create/attach publisher.")
		}
	}
	return nil
}

func createOrAttachPublisher(ctx context.Context, client *ent.Client, book *ent.Book, publisher calibre.Publisher) error {
	ic := importContextFrom(ctx)
	if ksuid, ok := ic.PublisherVisited(publisher.ID); ok {
		if err := book.Update().
			AddPublisherIDs(ksuid).
			Exec(ctx); err != nil {
			return fmt.Errorf("Error adding publisher to book: %w", err)
		}
		return nil
	}
	newPublisher, err := client.Publisher.Create().
		SetName(publisher.Name).
		SetCalibreID(publisher.ID).
		AddBooks(book).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create publisher (%d): %w", publisher.ID, err)
	}
	ic.AddPublisherVisited(publisher.ID, newPublisher.ID)
	return nil
}

func createOrAttachLanguages(ctx context.Context, client *ent.Client, book *ent.Book, languages []calibre.Language) error {
	for _, l := range languages {
		err := createOrAttachLanguage(ctx, client, book, l)
		if err != nil {
			log.Warn().Err(err).
				Str("language", l.LangCode).
				Msg("Failed to create/attach language.")
		}
	}
	return nil
}

func createOrAttachLanguage(ctx context.Context, client *ent.Client, book *ent.Book, language calibre.Language) error {
	ic := importContextFrom(ctx)
	if ksuid, ok := ic.LanguageVisited(language.ID); ok {
		if err := book.Update().
			AddLanguageIDs(ksuid).
			Exec(ctx); err != nil {
			return fmt.Errorf("Error adding language to book: %w", err)
		}
		return nil
	}
	newLanguage, err := client.Language.Create().
		SetCode(language.LangCode).
		SetCalibreID(language.ID).
		AddBooks(book).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create language (%d): %w", language.ID, err)
	}
	ic.AddLanguageVisited(language.ID, newLanguage.ID)
	return nil
}

func createOrAttachSeriesList(ctx context.Context, client *ent.Client, book *ent.Book, series []calibre.Series) error {
	for _, s := range series {
		err := createOrAttachSeries(ctx, client, book, s)
		if err != nil {
			log.Warn().Err(err).
				Str("series", s.Name).
				Msg("Failed to create/attach series.")
		}
	}
	return nil
}

func createOrAttachSeries(ctx context.Context, client *ent.Client, book *ent.Book, series calibre.Series) error {
	ic := importContextFrom(ctx)
	if ksuid, ok := ic.SeriesVisited(series.ID); ok {
		if err := book.Update().
			AddSeriesIDs(ksuid).
			Exec(ctx); err != nil {
			return fmt.Errorf("Error adding series to book: %w", err)
		}
		return nil
	}
	newSeries, err := client.Series.Create().
		SetName(series.Name).
		SetSort(series.Sort).
		SetCalibreID(series.ID).
		AddBooks(book).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create series (%d): %w", series.ID, err)
	}
	ic.AddSeriesVisited(series.ID, newSeries.ID)
	return nil
}

func registerBookFiles(ctx context.Context, cal calibre.Calibre, client *ent.Client, book *ent.Book, calBook calibre.Book) error {
	log := log.Ctx(ctx).With().Str("book", calBook.Title).Str("bookID", book.ID.String()).Logger()
	path := calBook.FullPath(cal)
	log.Info().Str("path", path).Msg("Registering book files")
	files, err := os.ReadDir(calBook.FullPath(cal))
	if err != nil {
		return fmt.Errorf("failed to read book directory: %w", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(calBook.FullPath(cal), file.Name())
		err := registerBookFile(ctx, client, book.ID, path, file)
		if err != nil {
			log.Warn().Err(err).
				Str("file", file.Name()).
				Msg("Failed to register file")
		}
	}
	return nil
}

func registerBookFile(ctx context.Context, client *ent.Client, bookID ksuid.ID, path string, file fs.DirEntry) error {
	ext := strings.ToLower(filepath.Ext(file.Name()))
	nameWithoutExt := strings.TrimSuffix(file.Name(), ext)
	if slices.Contains(ignorableFilenames, nameWithoutExt) {
		return nil
	}

	filetype := filetype.FromExtension(ext)
	if filetype == 0 {
		return errors.New("unknown file type")
	}
	info, err := file.Info()
	if err != nil {
		log.Warn().Err(err).
			Str("file", file.Name()).
			Msg("Failed to get file info")
		return fmt.Errorf("failed to get file info: %w", err)
	}
	size := info.Size()
	_, err = client.BookFile.Create().
		SetName(file.Name()).
		SetPath(path).
		SetSize(size).
		SetFormat(bookfile.Format(filetype.String())).
		SetBookID(bookID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create book file: %w", err)
	}
	return nil
}

// TODO: This really should be as simple as a bulk upsert, but this is blocked by https://github.com/ent/ent/issues/3868
// Example of how to do bulk upserts post https://github.com/ent/ent/issues/3868
// func createOrAttachTags(client *ent.Client, ctx context.Context, book *ent.Book, tags []calibre.Tag) error {
// 	tagCreates := make([]*ent.TagCreate, len(tags))
// 	for _, t := range tags {
// 		tgc := client.Tag.Create().
// 			SetName(t.Name).
// 			AddBooks(book)
// 		tagCreates = append(tagCreates, tgc)
// 	}
// 	return client.Tag.
// 		CreateBulk(tagCreates...).
// 		OnConflictColumns(tag.FieldName).
// 		UpdateNewValues().
// 		Exec(ctx)
// }
