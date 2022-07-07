# File System

This package, **filesystem**, contains interfaces and structures for abstracting and wrapping file system interactions. Use this in lieu of the direct `os` and `ioutil` packages for manipulating files on a physical file system.

## Interfaces

This package provides five interfaces which abstract a file system, a file object, and file reader/writer. 

### FileSystem

The interface **FileSystem** represents a system that reads and writes files and directories. It implements the following interfaces: [fs.FS](https://pkg.go.dev/io/fs@go1.18.3#FS), [fs.ReadFileFS](https://pkg.go.dev/io/fs@go1.18.3#ReadFileFS), **OpenFileFS**, **WriteFileFS**, and **DirWriterFS**.

This means that any struct that implements **FileSystem** must provide the following methods:

* Mkdir(name string, perm fs.FileMode) error
* Open(name string) ([fs.File](https://pkg.go.dev/io/fs@go1.18.3#File), error)
* OpenFile(name string, flag int, perm os.FileMode) (WriteableFile, error)
* ReadFile(name string) ([]byte, error)
* Stat(name string) ([fs.FileInfo](https://pkg.go.dev/io/fs@go1.18.3#FileInfo), error)
* WriteFile(name string, data []byte, perm fs.FileMode) error

Use this interface when you need the full ability to read and write files.

### DirWriterFS

**DirWriterFS** describes an interface to write directories in a file system. It provides the following:

* Mkdir(name string, perm fs.FileMode) error

### OpenFileFS

**OpenFileFS** describes an interface to open a file and specify what kind of action you wish to perform on this file, such as reading, writing, and/or appending. This is similar to the Go [fs.FS](https://pkg.go.dev/io/fs@go1.18.3#FS) interface but adds the flag attribute.

* OpenFile(name string, flag int, perm os.FileMode) (WriteableFile, error)

### WriteFileFS

**WriteFileFS** describes an interface to write bytes to a file system. 

* WriteFile(name string, data []byte, perm fs.FileMode) error

### WritableFile

**WritableFile** describes an interface to a file that can be both read and written to. This interface adds in [fs.File](https://pkg.go.dev/io/fs@go1.18.3#File) for the reading interface, then adds additional methods for writing.

* Close() error
* Name() string
* Read(b []byte) (int, error)
* Stat() ([fs.FileInfo](https://pkg.go.dev/io/fs@go1.18.3#FileInfo), error)
* Write(b []byte) (int, error)
* WriteString(s string) (int, error)

## Implementation

This package has the following packages that implement the interfaces described above.

* localfs
* memoryfs

### localfs

**localfs** provides structs that implement the file system interfaces to work directly with a local OS file system. It implements the interface **FileSystem**.

```go
package main

import (
	"fmt"

	"github.com/app-nerds/kit/v6/filesystem/localfs"
)

func main() {
	localFS := localfs.NewLocalFS()
	fileName := "test.txt"
	data := []byte("This is a text file!")

	if err := writeData(localFS, fileName, data); err != nil {
		panic("oh no!")
	}

	fmt.Printf("file written")
}

func writeData(fs filesystem.FileSystem, fileName string, data []byte) error {
	return fs.WriteFile(name, data, os.ModeAppend)
}
```

### memoryfs

**memoryfs** provides structs that implement the file system interfaces to work with a file system that exists in memory. It implements the interface **FileSystem**.

```go
package main

import (
	"fmt"

	"github.com/app-nerds/kit/v6/filesystem/memoryfs"
)

func main() {
	memoryFS := memoryfs.NewMemoryFS()
	fileName := "test.txt"
	data := []byte("This is a text file!")

	if err := writeData(localFS, fileName, data); err != nil {
		panic("oh no!")
	}

	fmt.Printf("file written")
}

func writeData(fs filesystem.FileSystem, fileName string, data []byte) error {
	return fs.WriteFile(name, data, os.ModeAppend)
}
```

