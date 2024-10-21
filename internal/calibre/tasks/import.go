package tasks

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"lybbrio/internal/calibre"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/schema/filetype"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/image"
	"lybbrio/internal/scheduler"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
)

func ImportTask(cal calibre.Calibre, client *ent.Client) scheduler.TaskFunc {
	return func(ctx context.Context, task *ent.Task, cb scheduler.ProgressCallback) (string, error) {
		log := log.Ctx(ctx)
		log.Info().Interface("task", task.ID.String()).Msg("ImportTask")

		ic := newImportContext()
		ctx = importContextTo(ctx, ic)

		err := importBooks(ctx, cal, client, cb)
		if err != nil {
			return "", err
		}

		return ic.String(), nil
	}
}

func importBooks(ctx context.Context, cal calibre.Calibre, client *ent.Client, cb scheduler.ProgressCallback) error {
	calibreBooks, err := cal.GetBooks(ctx)
	if err != nil {
		return err
	}
	ic := importContextFrom(ctx)
	total := len(calibreBooks)
	for idx, calBook := range calibreBooks {
		err := importBook(ctx, cal, client, *calBook)
		if err != nil {
			log.Warn().Err(err).
				Str("book", calBook.Title).
				Int64("bookID", calBook.ID).
				Msg("Failed to import book")
		}
		if err := cb(float64(idx+1) / (float64(total))); err != nil {
			ic.AddFailedBook(calBook.Title)
			log.Warn().Err(err).
				Str("book", calBook.Title).
				Int64("bookID", calBook.ID).
				Msg("Failed to update progress")
		}
	}

	return nil
}

func importBook(ctx context.Context, cal calibre.Calibre, client *ent.Client, calBook calibre.Book) error {

	bookCreate := client.Book.Create().
		SetTitle(calBook.Title).
		SetSort(calBook.Sort).
		SetCalibreID(calBook.ID).
		SetIsbn(calBook.ISBN).
		SetPath(calBook.Path).
		SetDescription(calBook.Comments.Text)
	if calBook.PubDate != nil {
		bookCreate.SetPublishedDate(*calBook.PubDate)
	}
	if calBook.SeriesIndex != nil {
		bookCreate.SetSeriesIndex(*calBook.SeriesIndex)
	}

	var entBook *ent.Book
	var err error
	entBook, err = bookCreate.
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			log.Debug().Err(err).
				Str("book", calBook.Title).
				Int64("bookID", calBook.ID).
				Msg("Book already exists")
			entBook, err = client.Book.Query().
				Where(book.CalibreIDEQ(int64(calBook.ID))).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("failed to query existing book: %w", err)
			}
		} else {
			return fmt.Errorf("failed to create book: %w", err)
		}
	}

	err = createOrAttachAuthors(ctx, client, entBook, calBook.Authors)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create authors")
	}

	err = createIdentifiers(ctx, client, entBook, calBook.Identifiers)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create identifiers")
	}

	err = createOrAttachTags(ctx, client, entBook, calBook.Tags)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create tags")
	}

	err = createOrAttachPublishers(ctx, client, entBook, calBook.Publisher)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create publishers")
	}

	err = createOrAttachLanguages(ctx, client, entBook, calBook.Languages)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create languages")
	}

	err = createOrAttachSeriesList(ctx, client, entBook, calBook.Series)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed to create series")
	}

	err = registerBookFiles(ctx, cal, client, entBook, calBook)
	if err != nil {
		log.Warn().Err(err).
			Str("book", calBook.Title).
			Int64("bookID", calBook.ID).
			Msg("Failed register book files")
	}
	if calBook.HasCover {
		err = registerBookCovers(ctx, cal, client, entBook, calBook)
		if err != nil {
			log.Warn().Err(err).
				Str("book", calBook.Title).
				Int64("bookID", calBook.ID).
				Msg("Failed register book covers")
		}
	}
	return nil
}

func createOrAttachAuthors(ctx context.Context, client *ent.Client, book *ent.Book, authors []calibre.Author) error {
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
	log.Trace().Str("path", path).Msg("Registering book files")
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
			if inner := errors.Unwrap(err); ent.IsConstraintError(inner) {
				continue
			}
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

	ft := filetype.FromExtension(ext)
	if ft == filetype.Unknown {
		return errors.New("unknown file type")
	}
	info, err := file.Info()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}
	size := info.Size()
	_, err = client.BookFile.Create().
		SetName(nameWithoutExt).
		SetPath(path).
		SetSize(size).
		SetFormat(bookfile.Format(ft.String())).
		SetBookID(bookID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create book file: %w", err)
	}
	return nil
}

func registerBookCovers(ctx context.Context, cal calibre.Calibre, client *ent.Client, book *ent.Book, calBook calibre.Book) error {
	log := log.Ctx(ctx).With().Str("book", calBook.Title).Str("bookID", book.ID.String()).Logger()
	path := calBook.FullPath(cal)
	log.Trace().Str("path", path).Msg("Registering book covers")
	files, err := os.ReadDir(calBook.FullPath(cal))
	if err != nil {
		return fmt.Errorf("failed to read book directory: %w", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(calBook.FullPath(cal), file.Name())
		err := registerBookCover(ctx, client, book.ID, path, file)
		if err != nil {
			if inner := errors.Unwrap(err); ent.IsConstraintError(inner) {
				continue
			}
			log.Warn().Err(err).
				Str("file", file.Name()).
				Msg("Failed to register cover")
		}
	}
	return nil
}

func registerBookCover(ctx context.Context, client *ent.Client, bookID ksuid.ID, path string, file fs.DirEntry) error {
	ext := strings.ToLower(filepath.Ext(file.Name()))
	nameWithoutExt := strings.TrimSuffix(file.Name(), ext)
	if !strings.HasPrefix(nameWithoutExt, "cover") {
		return nil
	}
	img, err := image.ProcessFile(path)
	if err != nil {
		return fmt.Errorf("failed to process image: %w", err)
	}

	_, err = client.BookCover.Create().
		SetPath(path).
		SetSize(img.Size).
		SetContentType(img.ContentType).
		SetWidth(img.Width).
		SetHeight(img.Height).
		SetBookID(bookID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create book cover: %w", err)
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
