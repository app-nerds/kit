package filesystem

import (
	"io/fs"
	"time"
)

type FileInfo struct {
	FileName     string
	FileSize     int64
	FileMode     fs.FileMode
	ModifiedTime time.Time
	IsDirectory  bool
	System       interface{}
}

func (fi FileInfo) Name() string {
	return fi.FileName
}

func (fi FileInfo) Size() int64 {
	return fi.FileSize
}

func (fi FileInfo) Mode() fs.FileMode {
	return fi.FileMode
}

func (fi FileInfo) ModTime() time.Time {
	return fi.ModifiedTime
}

func (fi FileInfo) IsDir() bool {
	return fi.IsDirectory
}

func (fi FileInfo) Sys() interface{} {
	return fi.System
}
