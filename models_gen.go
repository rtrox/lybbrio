// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package lybbrio

import (
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task_enums"
)

type CreateShelfInput struct {
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Public      *bool      `json:"public,omitempty"`
	BookIDs     []ksuid.ID `json:"bookIDs,omitempty"`
}

// CreateTaskInput is used for create Task object.
// Input was generated by ent.
type CreateTaskInput struct {
	Type task_enums.TaskType `json:"type"`
}
