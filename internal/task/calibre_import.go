package task

import (
	"context"
	"lybbrio/internal/ent"

	"github.com/rs/zerolog/log"
)

func CalibreImportTask(ctx context.Context, task *ent.Task, client *ent.Client) (msg string, err error) {
	log := log.Ctx(ctx)
	log.Info().Interface("task", task).Msg("CalibreImportTask")
	return "done", nil
}
