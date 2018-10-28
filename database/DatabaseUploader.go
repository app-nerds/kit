package database

import (
	"io"
)

/*
A DatabaseUploader defines an interface for uploading files to a database
*/
type DatabaseUploader interface {
	Upload(reader io.Reader, name, path string) (*DatabaseUploadResponse, error)
}
