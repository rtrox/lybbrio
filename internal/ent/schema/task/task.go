package task

import (
	"fmt"
	"io"
	"strconv"
)

type TaskType string

const (
	// TaskTypeNoOp is a task that does nothing.
	TypeNoOp          TaskType = "noop"
	TypeCalibreImport          = "calibre_import"
)

func (TaskType) Values() (kinds []string) {
	for _, s := range []TaskType{
		TypeNoOp,
		TypeCalibreImport,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

func (t TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(t)))
}

func (t *TaskType) UnmarshalGQL(v interface{}) error {
	return t.Scan(v)
}

func (t *TaskType) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("task: expected a value")
	}
	switch src := src.(type) {
	case string:
		*t = TaskType(src)
	case TaskType:
		*t = src
	default:
		return fmt.Errorf("task: unexpected type %T", src)
	}
	return nil
}

func (t TaskType) Value() (interface{}, error) {
	return string(t), nil
}

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning        = "running"
	StatusSuccess        = "success"
	StatusFailure        = "failure"
)

func (Status) Values() (kinds []string) {
	for _, s := range []Status{
		StatusPending,
		StatusRunning,
		StatusSuccess,
		StatusFailure,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

func (s Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(s)))
}

func (s *Status) UnmarshalGQL(v interface{}) error {
	return s.Scan(v)
}

func (s *Status) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("task: expected a value")
	}
	switch src := src.(type) {
	case string:
		*s = Status(src)
	case Status:
		*s = src
	default:
		return fmt.Errorf("task: unexpected type %T", src)
	}
	return nil
}

func (s Status) Value() (interface{}, error) {
	return string(s), nil
}
