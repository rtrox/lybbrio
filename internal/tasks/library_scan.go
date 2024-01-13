package tasks

import (
	"context"
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/filetype"
	"lybbrio/internal/scheduler"
	"os"
	"path/filepath"

	"github.com/pirmd/epub"
	"github.com/rs/zerolog/log"
)

func getBookFileCount(ctx context.Context, libraryPath string) (int64, error) {
	var ret int64
	err := filepath.WalkDir(libraryPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Could not walk library path: %w", err)
		}
		if d.IsDir() {
			return nil
		}
		fileType := filetype.FromExtension(d.Name())
		if fileType == filetype.Unknown {
			return nil
		}
		ret++
		return nil
	})
	return ret, err
}

func LibraryScanTask(client *ent.Client, libraryPath string) scheduler.TaskFunc {
	return func(ctx context.Context, task *ent.Task, cb scheduler.ProgressCallback) (string, error) {
		log := log.Ctx(ctx).With().Str("task_id", task.ID.String()).Logger()
		bookCount, err := getBookFileCount(ctx, libraryPath)
		if err != nil {
			return "", fmt.Errorf("Could not get book count: %w", err)
		}
		log.Info().Int64("book_count", bookCount).Msg("Found books in library")

		err = filepath.WalkDir(libraryPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("Could not walk library path: %w", err)
			}
			if d.IsDir() {
				return nil
			}
			fileType := filetype.FromExtension(d.Name())
			if fileType == filetype.Unknown {
				return nil
			}
			switch fileType {
			case filetype.EPUB:
				fallthrough
			case filetype.KEPUB:
				err = epubMetadata(ctx, path)
				if err != nil {
					return fmt.Errorf("Could not get epub metadata: %w", err)
				}
			}

			// TODO: Add book to database
			return nil
		})
		if err != nil {
			return "", fmt.Errorf("Could not walk library path: %w", err)
		}
		return "", nil
	}
}

func epubMetadata(ctx context.Context, path string) error {
	ep, err := epub.Open(path)
	if err != nil {
		return fmt.Errorf("Could not open epub: %w", err)
	}
	defer ep.Close()
	info, err := ep.Information()
	if err != nil {
		return fmt.Errorf("Could not get epub information: %w", err)
	}
	log := log.Ctx(ctx).With().Str("title", info.Title[0]).Logger()
	log.Info().Msg("Found epub")
	return nil
}
