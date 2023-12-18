package ksuid

import (
	"fmt"
	"io"
	"strconv"

	ksuid1 "github.com/segmentio/ksuid"
)

type ID string

func newKSUID() string {
	return ksuid1.New().String()
}

func MustNew(prefix string) ID {
	return ID(prefix + "_" + newKSUID())
}

func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

func (u ID) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(u)))
}

func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("ksuid: expected a value")
	}
	switch src := src.(type) {
	case string:
		*u = ID(src)
	case ID:
		*u = src
	default:
		return fmt.Errorf("ksuid: unexpected type %T", src)
	}
	return nil
}

func (u ID) Value() (interface{}, error) {
	return string(u), nil
}
