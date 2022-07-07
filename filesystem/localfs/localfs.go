package localfs

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/app-nerds/kit/v6/filesystem"
)

/*
LocalFile wraps an os.File struct. This implements WriableFile
*/
type LocalFile struct {
	f        os.File
	fileName string
	path     string
}

func (lf LocalFile) Stat() (fs.FileInfo, error) {
	return os.Stat(filepath.Join(lf.path, lf.fileName))
}

func (lf LocalFile) Close() error {
	return lf.f.Close()
}

func (lf LocalFile) Name() string {
	return filepath.Join(lf.path, lf.fileName)
}

func (lf LocalFile) Read(b []byte) (int, error) {
	return lf.f.Read(b)
}

func (lf LocalFile) Write(b []byte) (int, error) {
	return lf.f.Write(b)
}

func (lf LocalFile) WriteString(s string) (int, error) {
	return lf.f.WriteString(s)
}

/*
LocalFS is wrapper around the OS file system used to read and write files. This implements
fs.FS, fs.ReadFileFS, OpenFileFS, WriteFileFS.
*/
type LocalFS struct {
}

func NewLocalFS() filesystem.FileSystem {
	return &LocalFS{}
}

func (lfs *LocalFS) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(name, perm)
}

func (lfs *LocalFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (lfs *LocalFS) OpenFile(name string, flag int, perm os.FileMode) (filesystem.WritableFile, error) {
	return os.OpenFile(name, flag, perm)
}

func (lfs *LocalFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (lfs *LocalFS) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (lfs *LocalFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (lfs *LocalFS) ReadDir(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

func (lfs *LocalFS) FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
