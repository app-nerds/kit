package database

import (
	"io"
	"path/filepath"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

/*
Dial establishes a new session to one or more Mongo databases
*/
func Dial(url string) (Session, error) {
	s, err := mgo.Dial(url)

	result := &MongoSession{
		Session: s,
	}

	return result, err
}

// Session is an interface to access to the Session struct.
type Session interface {
	DB(name string) Database
	Close()
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) Database {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// DataLayer is an interface to access to the database struct.
type Database interface {
	C(name string) Collection
	GridFS(prefix string) GridFS
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

func (d MongoDatabase) GridFS(prefix string) GridFS {
	return &MongoGridFS{GridFS: d.Database.GridFS(prefix)}
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Count() (int, error)
	DropAllIndexes() error
	DropCollection() error
	DropIndex(key ...string) error
	DropIndexName(name string) error
	EnsureIndex(index mgo.Index) error
	EnsureIndexKey(key ...string) error
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Indexes() ([]mgo.Index, error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (*mgo.ChangeInfo, error)
	RemoveId(id interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	UpsertId(id interface{}, update interface{}) (*mgo.ChangeInfo, error)
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

func (c *MongoCollection) Count() (int, error) {
	return c.Collection.Count()
}

func (c *MongoCollection) DropAllIndexes() error {
	return c.Collection.DropAllIndexes()
}

func (c *MongoCollection) DropCollection() error {
	return c.Collection.DropCollection()
}

func (c *MongoCollection) DropIndex(key ...string) error {
	return c.Collection.DropIndex(key...)
}

func (c *MongoCollection) DropIndexName(name string) error {
	return c.DropIndexName(name)
}

func (c *MongoCollection) EnsureIndex(index mgo.Index) error {
	return c.Collection.EnsureIndex(index)
}

func (c *MongoCollection) EnsureIndexKey(key ...string) error {
	return c.Collection.EnsureIndexKey(key...)
}

// Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Find(query interface{}) Query {
	return &MongoQuery{Query: c.Collection.Find(query)}
}

func (c *MongoCollection) FindId(id interface{}) Query {
	return &MongoQuery{Query: c.Collection.FindId(id)}
}

func (c *MongoCollection) Indexes() ([]mgo.Index, error) {
	return c.Collection.Indexes()
}

func (c *MongoCollection) Insert(docs ...interface{}) error {
	return c.Collection.Insert(docs...)
}

func (c *MongoCollection) Remove(selector interface{}) error {
	return c.Collection.Remove(selector)
}

func (c *MongoCollection) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	return c.Collection.RemoveAll(selector)
}

func (c *MongoCollection) RemoveId(id interface{}) error {
	return c.Collection.RemoveId(id)
}

func (c *MongoCollection) Update(selector interface{}, update interface{}) error {
	return c.Collection.Update(selector, update)
}

func (c *MongoCollection) UpdateId(id interface{}, update interface{}) error {
	return c.Collection.UpdateId(id, update)
}

func (c *MongoCollection) Upsert(selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.Collection.Upsert(selector, update)
}

func (c *MongoCollection) UpsertId(id interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.Collection.UpsertId(id, update)
}

// Query is an interface to access to the database struct
type Query interface {
	All(result interface{}) error
	Count() (int, error)
	Distinct(key string, result interface{}) error
	Limit(n int) Query
	One(result interface{}) (err error)
	Select(selector interface{}) Query
	Skip(n int) Query
	Sort(fields ...string) Query
}

// MongoQuery wraps a mgo.Query to embed methods in models.
type MongoQuery struct {
	*mgo.Query
}

func (q *MongoQuery) All(result interface{}) error {
	return q.Query.All(result)
}

func (q *MongoQuery) Count() (int, error) {
	return q.Query.Count()
}

func (q *MongoQuery) Distinct(key string, result interface{}) error {
	return q.Query.Distinct(key, result)
}

func (q *MongoQuery) Limit(n int) Query {
	result := q.Query.Limit(n)
	return &MongoQuery{Query: result}
}

func (q *MongoQuery) One(result interface{}) error {
	return q.Query.One(result)
}

func (q *MongoQuery) Select(selector interface{}) Query {
	result := q.Query.Select(selector)
	return &MongoQuery{Query: result}
}

func (q *MongoQuery) Skip(n int) Query {
	result := q.Query.Skip(n)
	return &MongoQuery{Query: result}
}

func (q *MongoQuery) Sort(fields ...string) Query {
	result := q.Query.Sort(fields...)
	return &MongoQuery{Query: result}
}

// GridFS stores files in a MongoDB database
type GridFS interface {
	Create(name string) (GridFile, error)
	Find(query interface{}) Query
	Open(name string) (GridFile, error)
	OpenId(id interface{}) (GridFile, error)
	Remove(name string) (err error)
	RemoveId(id interface{}) error
}

// MongoGridFS is the wrapper around GridFS
type MongoGridFS struct {
	*mgo.GridFS
}

// Create makes a new GridFS storage
func (gfs *MongoGridFS) Create(name string) (GridFile, error) {
	result, err := gfs.GridFS.Create(name)
	return &MongoGridFile{GridFile: result}, err
}

// Find queries GridFS
func (gfs *MongoGridFS) Find(query interface{}) Query {
	return &MongoQuery{Query: gfs.GridFS.Find(query)}
}

// Open opens a grid file
func (gfs *MongoGridFS) Open(name string) (GridFile, error) {
	result, err := gfs.GridFS.Open(name)
	return &MongoGridFile{GridFile: result}, err
}

func (gfs *MongoGridFS) OpenId(id interface{}) (GridFile, error) {
	result, err := gfs.GridFS.OpenId(id)
	return &MongoGridFile{GridFile: result}, err
}

func (gfs *MongoGridFS) Remove(name string) error {
	return gfs.GridFS.Remove(name)
}
func (gfs *MongoGridFS) RemoveId(id interface{}) error {
	return gfs.RemoveId(id)
}

// GridFile represents a single file in GridFS
type GridFile interface {
	Abort()
	Close() error
	ContentType() string
	GetMeta(result interface{}) error
	Id() interface{}
	MD5() string
	Name() string
	Read(b []byte) (int, error)
	Seek(offset int64, whence int) (int64, error)
	SetChunkSize(bytes int)
	SetContentType(ctype string)
	SetId(id interface{})
	SetMeta(metadata interface{})
	SetName(name string)
	SetUploadDate(t time.Time)
	Size() int64
	Write(data []byte) (int, error)
}

type MongoGridFile struct {
	*mgo.GridFile
}

func (mgf *MongoGridFile) Abort() {
	mgf.GridFile.Abort()
}

func (mgf *MongoGridFile) Close() error {
	return mgf.GridFile.Close()
}

func (mgf *MongoGridFile) ContentType() string {
	return mgf.GridFile.ContentType()
}

func (mgf *MongoGridFile) GetMeta(result interface{}) error {
	return mgf.GridFile.GetMeta(result)
}

func (mgf *MongoGridFile) Id() interface{} {
	return mgf.GridFile.Id()
}

func (mgf *MongoGridFile) MD5() string {
	return mgf.GridFile.MD5()
}

func (mgf *MongoGridFile) Name() string {
	return mgf.GridFile.Name()
}

func (mgf *MongoGridFile) Read(b []byte) (int, error) {
	return mgf.GridFile.Read(b)
}

func (mgf *MongoGridFile) Seek(offset int64, whence int) (int64, error) {
	return mgf.GridFile.Seek(offset, whence)
}

func (mgf *MongoGridFile) SetChunkSize(bytes int) {
	mgf.GridFile.SetChunkSize(bytes)
}

func (mgf *MongoGridFile) SetContentType(ctype string) {
	mgf.GridFile.SetContentType(ctype)
}

func (mgf *MongoGridFile) SetId(id interface{}) {
	mgf.GridFile.SetId(id)
}

func (mgf *MongoGridFile) SetMeta(metadata interface{}) {
	mgf.GridFile.SetMeta(metadata)
}

func (mgf *MongoGridFile) SetName(name string) {
	mgf.GridFile.SetName(name)
}

func (mgf *MongoGridFile) SetUploadDate(t time.Time) {
	mgf.GridFile.SetUploadDate(t)
}

func (mgf *MongoGridFile) Size() int64 {
	return mgf.GridFile.Size()
}

func (mgf *MongoGridFile) Write(data []byte) (int, error) {
	return mgf.GridFile.Write(data)
}

type Iter interface {
	All(result interface{}) error
	Close() error
	Done() bool
	Err() error
	For(result interface{}, f func() error) error
	Next(result interface{}) bool
	State() (int64, []bson.Raw)
	Timeout() bool
}

type MongoIter struct {
	*mgo.Iter
}

func (iter *MongoIter) All(result interface{}) error {
	return iter.Iter.All(result)
}

func (iter *MongoIter) Close() error {
	return iter.Iter.Close()
}

func (iter *MongoIter) Done() bool {
	return iter.Iter.Done()
}

func (iter *MongoIter) Err() error {
	return iter.Iter.Err()
}

func (iter *MongoIter) For(result interface{}, f func() error) error {
	return iter.Iter.For(result, f)
}

func (iter *MongoIter) Next(result interface{}) bool {
	return iter.Iter.Next(result)
}

func (iter *MongoIter) State() (int64, []bson.Raw) {
	return iter.Iter.State()
}

func (iter *MongoIter) Timeout() bool {
	return iter.Iter.Timeout()
}

/*
A DatabaseUploader defines an interface for uploading files to a database
*/
type DatabaseUploader interface {
	Upload(reader io.Reader, name, path string) (*DatabaseUploadResponse, error)
}

/*
A DatabaseUploadResponse is used to alert a caller information about their uploaded file.
*/
type DatabaseUploadResponse struct {
	BytesWritten int           `json:"bytesWritten"`
	FileID       bson.ObjectId `json:"fileID"`
	FileName     string        `json:"fileName"`
	Height       int           `json:"height"`
	Width        int           `json:"width"`
}

/*
A MongoUploader uploads files to a MongoDB database. This struct
satisfies the DatabaseUploader interface
*/
type MongoUploader struct {
	DB Database
}

/*
Upload uploads a file into the MongoDB GridFS system
*/
func (u *MongoUploader) Upload(reader io.Reader, name, path string) (*DatabaseUploadResponse, error) {
	var err error
	var file GridFile

	result := &DatabaseUploadResponse{}
	totalBytesWritten := 0
	bytesWritten := 0
	bytesRead := 0

	buffer := make([]byte, 2048)
	name = u.sanitizeFileName(name)

	/*
	 * Create the file in GridFS
	 */
	if file, err = u.DB.GridFS(path).Create(name); err != nil {
		return result, errors.Wrapf(err, "Error uploading file '%s' to MongoDB GridFS", name)
	}

	defer file.Close()

	/*
	 * Read bytes from the file, write bytes to GridFS
	 */
	for {
		bytesRead, err = reader.Read(buffer)

		if err == nil {
			if bytesRead > 0 {
				if bytesWritten, err = file.Write(buffer); err != nil {
					result.BytesWritten = bytesWritten
					return result, errors.Wrapf(err, "Error writing file bytes to GridFS")
				}

				totalBytesWritten += bytesWritten
			}
		}

		if err == io.EOF {
			result.BytesWritten = totalBytesWritten
			result.FileID = (file.Id()).(bson.ObjectId)
			result.FileName = name

			return result, nil
		}
	}
}

func (u *MongoUploader) sanitizeFileName(name string) string {
	return filepath.Base(filepath.Clean(name))
}
