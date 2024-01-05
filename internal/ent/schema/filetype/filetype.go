package filetype

import "strings"

type FileType int

const (
	Unknown FileType = iota
	AZW3
	EPUB
	KEPUB
	PDF
	CBC
	CBR
	CB7
	CBZ
	CBT

	count // Keep this at the end.
)

func (f FileType) String() string {
	switch f {
	case AZW3:
		return "AZW3"
	case EPUB:
		return "EPUB"
	case KEPUB:
		return "KEPUB"
	case PDF:
		return "PDF"
	case CBC:
		return "CBC"
	case CBR:
		return "CBR"
	case CB7:
		return "CB7"
	case CBZ:
		return "CBZ"
	case CBT:
		return "CBT"
	}
	return ""
}

func FromString(s string) FileType {
	switch s {
	case AZW3.String():
		return AZW3
	case EPUB.String():
		return EPUB
	case KEPUB.String():
		return KEPUB
	case PDF.String():
		return PDF
	case CBC.String():
		return CBC
	case CBR.String():
		return CBR
	case CB7.String():
		return CB7
	case CBZ.String():
		return CBZ
	case CBT.String():
		return CBT
	default:
		return 0
	}
}

func FromExtension(s string) FileType {
	s = strings.Replace(strings.ToUpper(s), ".", "", 1)
	return FromString(s)
}

func All() []FileType {
	ret := make([]FileType, 0, count)
	for i := AZW3; i < count; i++ {
		ret = append(ret, i)
	}
	return ret
}
