/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package images

import (
	"bytes"
	"io"
)

type MockResizer struct {
	ResizeImageFunc       func(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error)
	ResizeImagePixelsFunc func(source io.ReadSeeker, contentType string, width, height int) (*bytes.Buffer, error)
}

func (m MockResizer) ResizeImage(source io.ReadSeeker, contentType string, imageSize ImageSize) (*bytes.Buffer, error) {
	return m.ResizeImageFunc(source, contentType, imageSize)
}

func (m MockResizer) ResizeImagePixels(source io.ReadSeeker, contentType string, width, height int) (*bytes.Buffer, error) {
	return m.ResizeImagePixelsFunc(source, contentType, width, height)
}
