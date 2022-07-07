package filesystem

import (
	"io/fs"
	"os"
)

type DirWriterFS interface {
	Mkdir(name string, perm fs.FileMode) error
}

type OpenFileFS interface {
	OpenFile(name string, flag int, perm os.FileMode) (WritableFile, error)
}

type WriteFileFS interface {
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

type ReadDirFS interface {
	ReadDir(dir string) ([]fs.DirEntry, error)
}

type FileExistsFS interface {
	FileExists(file string) bool
}

type WritableFile interface {
	fs.File

	Name() string
	Write(b []byte) (int, error)
	WriteString(s string) (int, error)
}

type FileSystem interface {
	fs.FS
	fs.ReadFileFS
	DirWriterFS
	OpenFileFS
	WriteFileFS
	ReadDirFS
	FileExistsFS

	Stat(name string) (fs.FileInfo, error)
}
