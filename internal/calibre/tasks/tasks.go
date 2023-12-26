package tasks

import (
	"lybbrio/internal/calibre"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/task"
)

func TaskMap(cal calibre.Calibre, client *ent.Client) task.TaskMap {
	return task.TaskMap{
		task_enums.TypeCalibreImport: ImportTask(cal, client),
	}
}
