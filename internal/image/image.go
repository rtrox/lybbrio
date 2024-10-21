package image

import (
	"bytes"
	"errors"
	img "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
)

var (
	ErrFileNotFound        = errors.New("file not found")
	ErrUnknownImageType    = errors.New("unknown image type")
	ErrorReadingFileHeader = errors.New("error reading file header")
)

type ImageFile struct {
	Size        int64
	Width       int
	Height      int
	ContentType string
}

func ProcessFile(filePath string) (ImageFile, error) {
	ret := ImageFile{}
	fi, err := os.Stat(filePath)
	if err != nil {
		return ImageFile{}, ErrFileNotFound
	}
	ret.Size = fi.Size()

	r, err := os.Open(filePath)
	if err != nil {
		// shouldn't be reachable, stat will fail if the file doesn't exist.
		return ImageFile{}, ErrFileNotFound
	}
	defer r.Close()

	im, _, err := img.DecodeConfig(r)
	if err != nil {
		return ImageFile{}, ErrUnknownImageType
	}

	ret.Width = im.Width
	ret.Height = im.Height

	if _, err := r.Seek(0, 0); err != nil {
		return ImageFile{}, ErrorReadingFileHeader
	}
	// Limit the reader to the first 512 bytes to read the file header.
	lr := io.LimitReader(r, 512)
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, lr)
	if err != nil {
		return ImageFile{}, ErrorReadingFileHeader
	}
	ret.ContentType = http.DetectContentType(b.Bytes())

	return ret, nil
}
