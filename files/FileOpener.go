/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package files

import (
	"os"
)

type FileOpener interface {
	Open(name string) (*os.File, error)
}

type MockFileOpener struct {
	OpenFunc func(name string) (*os.File, error)
}

func (m *MockFileOpener) Open(name string) (*os.File, error) {
	return m.OpenFunc(name)
}
