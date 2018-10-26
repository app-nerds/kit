package database

import "time"

type MongoGridFile interface {
	Abort()
	Close() error
	ContentType() string
	GetMeta(result interface{}) error
	Id() interface{}
	MD5() (md5 string)
	Name() string
	Read(b []byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
	SetChunkSize(bytes int)
	SetContentType(ctype string)
	SetId(id interface{})
	SetMeta(metadata interface{})
	SetName(name string)
	Size() (bytes int64)
	UploadDate() time.Time
	Write(data []byte) (int, error)
}
