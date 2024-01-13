package tasks

import (
	"lybbrio/internal/calibre"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/scheduler"
)

func TaskMap(cal calibre.Calibre, client *ent.Client) scheduler.TaskMap {
	return scheduler.TaskMap{
		task_enums.TypeCalibreImport: ImportTask(cal, client),
	}
}
