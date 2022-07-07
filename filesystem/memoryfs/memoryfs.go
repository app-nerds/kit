package memoryfs

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/app-nerds/kit/v6/filesystem"
)

type DirEntry struct {
	name    string
	dirType fs.FileMode
	info    fs.FileInfo
}

func (d *DirEntry) Name() string      { return d.name }
func (d *DirEntry) IsDir() bool       { return true }
func (d *DirEntry) Type() fs.FileMode { return d.dirType }

func (d *DirEntry) Info() (fs.FileInfo, error) {
	if d.info != nil {
		return d.info, nil
	}
	return d.info, nil
}

func NewDirEntry(name string) fs.DirEntry {
	return &DirEntry{
		name: name,
	}
}

/*
MemoryFile is an in-memory file struct that implements fs.File, WritableFile
*/
type MemoryFile struct {
	data     []byte
	fileName string
	path     string
	flag     int
}

func (mf *MemoryFile) Close() error {
	return nil
}

func (mf *MemoryFile) Name() string {
	return filepath.Join(mf.path, mf.fileName)
}

func (mf *MemoryFile) Read(b []byte) (int, error) {
	buffer := bytes.NewBuffer(b)
	return buffer.Write(mf.data)
}

func (mf *MemoryFile) Stat() (fs.FileInfo, error) {
	return filesystem.FileInfo{
		FileName:     mf.fileName,
		FileSize:     int64(len(mf.data)),
		FileMode:     0777,
		ModifiedTime: time.Now().UTC(),
		IsDirectory:  false,
		System:       nil,
	}, nil
}

func (mf *MemoryFile) Write(b []byte) (int, error) {
	if mf.flag&os.O_APPEND > 0 {
		newBuffer := bytes.NewBuffer(mf.data)
		_, _ = newBuffer.Write(b)

		mf.data = newBuffer.Bytes()
		return len(b), nil
	}

	mf.data = b
	return len(b), nil
}

func (mf *MemoryFile) WriteString(s string) (int, error) {
	if mf.flag&os.O_APPEND > 0 {
		newBuffer := bytes.NewBuffer(mf.data)
		_, _ = newBuffer.WriteString(s)

		mf.data = newBuffer.Bytes()
		return len(s), nil
	}

	mf.data = []byte(s)
	return len(s), nil
}

/*
MemoryFS is an in-memory file system used to read and write files. This implements
fs.FS, fs.ReadFileFS, OpenFileFS, WriteFileFS.
*/
type MemoryFS struct {
	dirs  []*DirEntry
	files []*MemoryFile
}

func NewMemoryFS() *MemoryFS {
	return &MemoryFS{
		dirs:  []*DirEntry{},
		files: []*MemoryFile{},
	}
}

func (mfs *MemoryFS) Mkdir(name string, perm fs.FileMode) error {
	mfs.dirs = append(mfs.dirs, &DirEntry{
		name:    name,
		dirType: perm,
	})

	return nil
}

func (mfs *MemoryFS) Open(name string) (fs.File, error) {
	for _, f := range mfs.files {
		if filepath.Join(f.path, f.fileName) == name {
			return f, nil
		}
	}

	return nil, os.ErrNotExist
}

func (mfs *MemoryFS) OpenFile(name string, flag int, perm os.FileMode) (filesystem.WritableFile, error) {
	var f *MemoryFile
	found := false

	for _, f = range mfs.files {
		if f.path == name {
			found = true
			break
		}
	}

	if !found && flag&os.O_CREATE == 0 && flag&os.O_APPEND == 0 {
		return f, os.ErrNotExist
	}

	if (!found && flag&os.O_APPEND > 0) || (!found && flag&os.O_CREATE > 0) {
		newFile := &MemoryFile{
			data:     []byte{},
			fileName: filepath.Base(name),
			path:     filepath.Dir(name),
			flag:     flag,
		}

		mfs.files = append(mfs.files, newFile)
		return newFile, nil
	}

	return f, nil
}

func (mfs *MemoryFS) ReadFile(name string) ([]byte, error) {
	var (
		err error
		f   fs.File
	)

	if f, err = mfs.Open(name); err != nil {
		return []byte{}, err
	}

	result := f.(*MemoryFile)
	return result.data, nil
}

func (mfs *MemoryFS) Stat(name string) (fs.FileInfo, error) {
	f, err := mfs.Open(name)

	if err != nil {
		return filesystem.FileInfo{}, err
	}

	file := f.(*MemoryFile)

	return filesystem.FileInfo{
		FileName:     name,
		FileSize:     int64(len(file.data)),
		FileMode:     0,
		ModifiedTime: time.Now(),
		IsDirectory:  false,
		System:       nil,
	}, nil
}

func (mfs *MemoryFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	newFile := &MemoryFile{
		data:     data,
		fileName: filepath.Base(name),
		path:     filepath.Dir(name),
		flag:     os.O_CREATE,
	}

	mfs.files = append(mfs.files, newFile)
	return nil
}

func (mfs *MemoryFS) ReadDir(dir string) ([]fs.DirEntry, error) {
	var (
		err      error
		result   []fs.DirEntry
		foundDir *DirEntry
	)

	for _, d := range mfs.dirs {
		if d.name == dir {
			foundDir = d
			break
		}
	}

	if foundDir == nil {
		return result, fmt.Errorf("directory not found")
	}

	for _, f := range mfs.files {
		fi := filesystem.FileInfo{
			FileName:     f.Name(),
			FileSize:     0,
			FileMode:     0777,
			ModifiedTime: time.Now().UTC(),
			IsDirectory:  false,
			System:       nil,
		}

		de := &DirEntry{
			name: f.Name(),
			info: fi,
		}

		result = append(result, de)
	}

	return result, nil
}

func (mfs *MemoryFS) FileExists(file string) bool {
	for _, f := range mfs.files {
		if f.Name() == file {
			return true
		}
	}

	return false
}
