package tasks

import (
	"context"
	"fmt"
	"lybbrio/internal/calibre"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/author"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/series"
	"lybbrio/internal/ent/tag"
	"lybbrio/internal/task"
	"strings"

	"github.com/rs/zerolog/log"
)

type importOutput struct {
	FailedAuthors []string
	FailedBooks   []string
}

func (o *importOutput) String() string {
	var ret strings.Builder
	if len(o.FailedAuthors) > 0 {
		ret.WriteString("Failed Authors:\n")
		for _, author := range o.FailedAuthors {
			ret.WriteString(fmt.Sprintf("\t%s\n", author))
		}
	}
	if len(o.FailedBooks) > 0 {
		ret.WriteString("Failed Books:\n")
		for _, book := range o.FailedBooks {
			ret.WriteString(fmt.Sprintf("\t%s\n", book))
		}
	}
	return ret.String()
}

type importOutputCtxKey string

const importOutputKey importOutputCtxKey = "importOutput"

func importOutputFromContext(ctx context.Context) *importOutput {
	output := ctx.Value(importOutputKey)
	if output == nil {
		return nil
	}
	return output.(*importOutput)
}

func importOutputToContext(ctx context.Context, output *importOutput) context.Context {
	return context.WithValue(ctx, importOutputKey, output)
}

func ImportTask(cal calibre.Calibre, client *ent.Client) task.TaskFunc {
	return func(ctx context.Context, task *ent.Task, client *ent.Client) (string, error) {
		log := log.Ctx(ctx)
		log.Info().Interface("task", task.ID.String()).Msg("ImportTask")

		output := &importOutput{}
		ctx = importOutputToContext(ctx, output)

		err := ImportAuthors(cal, client, ctx)
		if err != nil {
			return "", err
		}

		err = ImportBooks(cal, client, ctx)
		if err != nil {
			return "", err
		}

		return output.String(), nil
	}
}

func ImportAuthors(cal calibre.Calibre, client *ent.Client, ctx context.Context) error {
	authors, err := cal.GetAuthors(context.Background())
	if err != nil {
		return err
	}
	for _, author := range authors {
		_, err := client.Author.Create().
			SetName(author.Name).
			SetSort(author.Sort).
			SetLink(author.Link).
			SetCalibreID(author.ID).
			Save(ctx)
		if err != nil {
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Str("author", author.Name).
					Int64("authorID", author.ID).
					Msg("Author already exists")
			} else {
				output := importOutputFromContext(ctx)
				output.FailedAuthors = append(output.FailedAuthors, author.Name)
			}
			continue
		}
	}
	return nil
}

func ImportBooks(cal calibre.Calibre, client *ent.Client, ctx context.Context) error {
	books, err := cal.GetBooks(context.Background())
	if err != nil {
		return err
	}
	for _, book := range books {

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

		authorCalibreIDs := make([]int64, len(book.Authors))
		for _, author := range book.Authors {
			authorCalibreIDs = append(authorCalibreIDs, author.ID)
		}
		authors, err := client.Author.Query().
			Where(author.CalibreIDIn(authorCalibreIDs...)).
			All(ctx)
		if err != nil {
			log.Warn().
				Err(err).
				Int64("bookID", book.ID).
				Str("book", book.Title).
				Msg("Failed to query authors")
			output := importOutputFromContext(ctx)
			output.FailedBooks = append(output.FailedBooks, book.Title)
			continue
		}
		bookCreate.AddAuthors(authors...)

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
				output := importOutputFromContext(ctx)
				output.FailedBooks = append(output.FailedBooks, book.Title)
			}
			continue
		}

		err = createIdentifiers(client, ctx, entBook, book.Identifiers)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create identifiers")
		}

		err = createOrAttachTags(client, ctx, entBook, book.Tags)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create tags")
		}

		err = createOrAttachPublishers(client, ctx, entBook, book.Publisher)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create publishers")
		}

		err = createOrAttachLanguages(client, ctx, entBook, book.Languages)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create languages")
		}

		err = createOrAttachSeries(client, ctx, entBook, book.Series)
		if err != nil {
			log.Warn().Err(err).
				Str("book", book.Title).
				Int64("bookID", book.ID).
				Msg("Failed to create series")
		}
	}
	return nil
}

func createIdentifiers(client *ent.Client, ctx context.Context, book *ent.Book, identifiers []calibre.Identifier) error {
	// TODO: This really should be as simple as a bulk upsert, but this is blocked by https://github.com/ent/ent/issues/3868
	for _, i := range identifiers {
		err := client.Identifier.Create().
			SetType(i.Type).
			SetValue(i.Val).
			SetBook(book).
			SetCalibreID(i.ID).
			Exec(ctx)
		if err != nil {
			log := log.With().
				Str("identifier", i.Val).
				Str("bookID", book.ID.String()).
				Logger()
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Msg("Identifier already exists")
				id, err := client.Identifier.Query().
					Where(identifier.CalibreID(i.ID)).
					Only(ctx)
				if err != nil {
					log.Error().Err(err).
						Msg("Failed to query identifier")
					continue
				}
				err = id.Update().
					SetBook(book).
					Exec(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to attach identifier to book")
					continue
				}
			} else {
				log.Warn().Err(err).
					Msg("Failed to create identifier")
			}
		}
	}
	return nil
}

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

func createOrAttachTags(client *ent.Client, ctx context.Context, book *ent.Book, tags []calibre.Tag) error {
	for _, t := range tags {
		err := client.Tag.Create().
			SetName(t.Name).
			SetCalibreID(t.ID).
			AddBooks(book).
			Exec(ctx)
		if err != nil {
			log := log.With().
				Str("tag", t.Name).
				Str("bookID", book.ID.String()).
				Logger()
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Msg("Tag already exists")
				tag, err := client.Tag.Query().
					Where(tag.Name(t.Name)).
					Only(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to query tag")
				}
				err = tag.Update().
					AddBooks(book).
					Exec(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to attach tag to book")
					continue
				}
			} else {
				log.Warn().Err(err).
					Msg("Failed to create tag")
			}
		}
	}
	return nil
}

func createOrAttachPublishers(client *ent.Client, ctx context.Context, book *ent.Book, publishers []calibre.Publisher) error {
	for _, p := range publishers {
		err := client.Publisher.Create().
			SetName(p.Name).
			SetCalibreID(p.ID).
			AddBooks(book).
			Exec(ctx)
		if err != nil {
			log := log.With().
				Str("publisher", p.Name).
				Str("bookID", book.ID.String()).
				Logger()
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Msg("Publisher already exists")
				publisher, err := client.Publisher.Query().
					Where(publisher.Name(p.Name)).
					Only(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to query publisher")
				}
				err = publisher.Update().
					AddBooks(book).
					Exec(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to attach publisher to book")
					continue
				}
			} else {
				log.Warn().Err(err).
					Msg("Failed to create publisher")
			}
		}
	}
	return nil
}

func createOrAttachLanguages(client *ent.Client, ctx context.Context, book *ent.Book, languages []calibre.Language) error {
	for _, l := range languages {
		err := client.Language.Create().
			SetCode(l.LangCode).
			SetCalibreID(l.ID).
			AddBooks(book).
			Exec(ctx)
		if err != nil {
			log := log.With().
				Str("language", l.LangCode).
				Str("bookID", book.ID.String()).
				Logger()
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Msg("Language already exists")
				language, err := client.Language.Query().
					Where(language.Code(l.LangCode)).
					Only(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to query language")
				}
				err = language.Update().
					AddBooks(book).
					Exec(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to attach language to book")
					continue
				}
			} else {
				log.Warn().Err(err).
					Msg("Failed to create language")
			}
		}
	}
	return nil
}

func createOrAttachSeries(client *ent.Client, ctx context.Context, book *ent.Book, calSeries []calibre.Series) error {
	for _, s := range calSeries {
		err := client.Series.Create().
			SetName(s.Name).
			SetSort(s.Sort).
			SetCalibreID(s.ID).
			AddBooks(book).
			Exec(ctx)
		if err != nil {
			log := log.With().
				Str("series", s.Name).
				Str("bookID", book.ID.String()).
				Logger()
			if ent.IsConstraintError(err) {
				log.Debug().Err(err).
					Msg("Series already exists")
				entSeries, err := client.Series.Query().
					Where(series.Name(s.Name)).
					Only(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to query series")
				}
				err = entSeries.Update().
					AddBooks(book).
					Exec(ctx)
				if err != nil {
					log.Warn().Err(err).
						Msg("Failed to attach series to book")
					continue
				}
			} else {
				log.Warn().Err(err).
					Msg("Failed to create series")
			}
		}
	}
	return nil
}
