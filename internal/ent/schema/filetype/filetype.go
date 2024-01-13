package filetype

import "strings"

type FileType int

const (
	Unknown FileType = iota
	EPUB
	KEPUB
	// AZW3
	// PDF
	// CBC
	// CBR
	// CB7
	// CBZ
	// CBT

	count // Keep this at the end.
)

func (f FileType) String() string {
	switch f {
	case EPUB:
		return "EPUB"
	case KEPUB:
		return "KEPUB"
	}
	return ""
}

func FromString(s string) FileType {
	switch s {
	case EPUB.String():
		return EPUB
	case KEPUB.String():
		return KEPUB
	default:
		return Unknown
	}
}

func FromExtension(s string) FileType {
	s = strings.Replace(strings.ToUpper(s), ".", "", 1)
	return FromString(s)
}

func All() []FileType {
	ret := make([]FileType, 0, count)
	for i := EPUB; i < count; i++ {
		ret = append(ret, i)
	}
	return ret
}
