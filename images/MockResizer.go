package images

import (
	"bytes"
	"io"
)

type MockResizer struct {
	ResizeImageFunc func(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error)
}

func (m *MockResizer) ResizeImage(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error) {
	return m.ResizeImageFunc(source, contentType, imageSize)
}
