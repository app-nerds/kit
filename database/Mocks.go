/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package database

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

/*
SessionMock is a mock of Session
*/
type SessionMock struct {
	DBFunc    func(name string) Database
	CloseFunc func()
}

/*
DB is a mock function
*/
func (s *SessionMock) DB(name string) Database {
	return s.DBFunc(name)
}

/*
Close is a mock function
*/
func (s *SessionMock) Close() {
	s.CloseFunc()
}

/*
DatabaseMock is a mock Database
*/
type DatabaseMock struct {
	CFunc      func(name string) Collection
	GridFSFunc func(prefix string) GridFS
}

/*
C is a mock function
*/
func (d *DatabaseMock) C(name string) Collection {
	return d.CFunc(name)
}

/*
GridFS is a mock function
*/
func (d *DatabaseMock) GridFS(prefix string) GridFS {
	return d.GridFSFunc(prefix)
}

/*
CollectionMock is a mock Collection
*/
type CollectionMock struct {
	CountFunc          func() (int, error)
	DropAllIndexesFunc func() error
	DropCollectionFunc func() error
	DropIndexFunc      func(key ...string) error
	DropIndexNameFunc  func(name string) error
	EnsureIndexFunc    func(index mgo.Index) error
	EnsureIndexKeyFunc func(key ...string) error
	FindFunc           func(query interface{}) Query
	FindIdFunc         func(id interface{}) Query
	FindWithPagingFunc func(query interface{}, skip, limit int) (Query, int, error)
	IndexesFunc        func() ([]mgo.Index, error)
	InsertFunc         func(docs ...interface{}) error
	RemoveFunc         func(selector interface{}) error
	RemoveAllFunc      func(selector interface{}) (*mgo.ChangeInfo, error)
	RemoveIdFunc       func(id interface{}) error
	UpdateFunc         func(selector interface{}, update interface{}) error
	UpdateAllFunc      func(selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	UpdateIdFunc       func(id interface{}, update interface{}) error
	UpsertFunc         func(selector interface{}, update interface{}) (*mgo.ChangeInfo, error)
	UpsertIdFunc       func(id interface{}, update interface{}) (*mgo.ChangeInfo, error)
}

/*
Count is a mock function
*/
func (c *CollectionMock) Count() (int, error) {
	return c.CountFunc()
}

/*
DropAllIndexes is a mock function
*/
func (c *CollectionMock) DropAllIndexes() error {
	return c.DropAllIndexesFunc()
}

/*
DropCollection is a mock function
*/
func (c *CollectionMock) DropCollection() error {
	return c.DropCollectionFunc()
}

/*
DropIndex is a mock function
*/
func (c *CollectionMock) DropIndex(key ...string) error {
	return c.DropIndexFunc(key...)
}

/*
DropIndexName is a mock function
*/
func (c *CollectionMock) DropIndexName(name string) error {
	return c.DropIndexNameFunc(name)
}

/*
EnsureIndex is a mock function
*/
func (c *CollectionMock) EnsureIndex(index mgo.Index) error {
	return c.EnsureIndexFunc(index)
}

/*
EnsureIndexKey is a mock function
*/
func (c *CollectionMock) EnsureIndexKey(key ...string) error {
	return c.EnsureIndexKeyFunc(key...)
}

/*
Find is a mock function
*/
func (c *CollectionMock) Find(query interface{}) Query {
	return c.FindFunc(query)
}

/*
FindId is a mock function
*/
func (c *CollectionMock) FindId(id interface{}) Query {
	return c.FindIdFunc(id)
}

/*
FindWithPaging is a mock function
*/
func (c *CollectionMock) FindWithPaging(query interface{}, skip, limit int) (Query, int, error) {
	return c.FindWithPagingFunc(query, skip, limit)
}

/*
Indexes is a mock function
*/
func (c *CollectionMock) Indexes() ([]mgo.Index, error) {
	return c.IndexesFunc()
}

/*
Insert is a mock function
*/
func (c *CollectionMock) Insert(docs ...interface{}) error {
	return c.InsertFunc(docs...)
}

/*
Remov is a mock functione
*/
func (c *CollectionMock) Remove(selector interface{}) error {
	return c.RemoveFunc(selector)
}

/*
RemoveAll is a mock function
*/
func (c *CollectionMock) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	return c.RemoveAllFunc(selector)
}

/*
RemoveId is a mock function
*/
func (c *CollectionMock) RemoveId(id interface{}) error {
	return c.RemoveIdFunc(id)
}

/*
Update is a mock function
*/
func (c *CollectionMock) Update(selector interface{}, update interface{}) error {
	return c.UpdateFunc(selector, update)
}

/*
UpdateAll is a mock function
*/
func (c *CollectionMock) UpdateAll(selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpdateAllFunc(selector, update)
}

/*
UpdateId is a mock function
*/
func (c *CollectionMock) UpdateId(id interface{}, update interface{}) error {
	return c.UpdateIdFunc(id, update)
}

/*
Upsert is a mock function
*/
func (c *CollectionMock) Upsert(selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertFunc(selector, update)
}

/*
UpsertId is a mock function
*/
func (c *CollectionMock) UpsertId(id interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertIdFunc(id, update)
}

/*
QueryMock mocks a Query
*/
type QueryMock struct {
	AllFunc      func(result interface{}) error
	CountFunc    func() (int, error)
	DistinctFunc func(key string, result interface{}) error
	LimitFunc    func(n int) Query
	OneFunc      func(result interface{}) error
	SelectFunc   func(selector interface{}) Query
	SkipFunc     func(n int) Query
	SortFunc     func(fields ...string) Query
}

/*
All is a mock function
*/
func (q *QueryMock) All(result interface{}) error {
	return q.AllFunc(result)
}

/*
All is a mock function
*/
func (q *QueryMock) Count() (int, error) {
	return q.CountFunc()
}

/*
All is a mock function
*/
func (q *QueryMock) Distinct(key string, result interface{}) error {
	return q.DistinctFunc(key, result)
}

/*
All is a mock function
*/
func (q *QueryMock) Limit(n int) Query {
	return q.LimitFunc(n)
}

/*
All is a mock function
*/
func (q *QueryMock) One(result interface{}) error {
	return q.OneFunc(result)
}

/*
All is a mock function
*/
func (q *QueryMock) Select(selector interface{}) Query {
	return q.SelectFunc(selector)
}

/*
All is a mock function
*/
func (q *QueryMock) Skip(n int) Query {
	return q.SkipFunc(n)
}

/*
All is a mock function
*/
func (q *QueryMock) Sort(fields ...string) Query {
	return q.SortFunc(fields...)
}

/*
GridFSMock mocks GridFS
*/
type GridFSMock struct {
	CreateFunc   func(name string) (GridFile, error)
	FindFunc     func(query interface{}) Query
	OpenFunc     func(name string) (GridFile, error)
	OpenIdFunc   func(id interface{}) (GridFile, error)
	RemoveFunc   func(name string) (err error)
	RemoveIdFunc func(id interface{}) error
}

/*
All is a mock function
*/
func (g *GridFSMock) Create(name string) (GridFile, error) {
	return g.CreateFunc(name)
}

/*
All is a mock function
*/
func (g *GridFSMock) Find(query interface{}) Query {
	return g.FindFunc(query)
}

/*
All is a mock function
*/
func (g *GridFSMock) Open(name string) (GridFile, error) {
	return g.OpenFunc(name)
}

/*
All is a mock function
*/
func (g *GridFSMock) OpenId(id interface{}) (GridFile, error) {
	return g.OpenIdFunc(id)
}

/*
All is a mock function
*/
func (g *GridFSMock) Remove(name string) (err error) {
	return g.RemoveFunc(name)
}

/*
All is a mock function
*/
func (g *GridFSMock) RemoveId(id interface{}) error {
	return g.RemoveIdFunc(id)
}

/*
GridFileMock is a mock GridFile
*/
type GridFileMock struct {
	AbortFunc          func()
	CloseFunc          func() error
	ContentTypeFunc    func() string
	GetMetaFunc        func(result interface{}) error
	IdFunc             func() interface{}
	MD5Func            func() string
	NameFunc           func() string
	ReadFunc           func(b []byte) (int, error)
	SeekFunc           func(offset int64, whence int) (int64, error)
	SetChunkSizeFunc   func(bytes int)
	SetContentTypeFunc func(ctype string)
	SetIdFunc          func(id interface{})
	SetMetaFunc        func(metadata interface{})
	SetNameFunc        func(name string)
	SetUploadDateFunc  func(t time.Time)
	SizeFunc           func() int64
	WriteFunc          func(data []byte) (int, error)
}

/*
All is a mock function
*/
func (g *GridFileMock) Abort() {
	g.AbortFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) Close() error {
	return g.CloseFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) ContentType() string {
	return g.ContentTypeFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) GetMeta(result interface{}) error {
	return g.GetMetaFunc(result)
}

/*
All is a mock function
*/
func (g *GridFileMock) Id() interface{} {
	return g.IdFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) MD5() string {
	return g.MD5Func()
}

/*
All is a mock function
*/
func (g *GridFileMock) Name() string {
	return g.NameFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) Read(b []byte) (int, error) {
	return g.ReadFunc(b)
}

/*
All is a mock function
*/
func (g *GridFileMock) Seek(offset int64, whence int) (int64, error) {
	return g.SeekFunc(offset, whence)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetChunkSize(bytes int) {
	g.SetChunkSizeFunc(bytes)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetContentType(ctype string) {
	g.SetContentTypeFunc(ctype)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetId(id interface{}) {
	g.SetIdFunc(id)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetMeta(metadata interface{}) {
	g.SetMetaFunc(metadata)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetName(name string) {
	g.SetName(name)
}

/*
All is a mock function
*/
func (g *GridFileMock) SetUploadDate(t time.Time) {
	g.SetUploadDateFunc(t)
}

/*
All is a mock function
*/
func (g *GridFileMock) Size() int64 {
	return g.SizeFunc()
}

/*
All is a mock function
*/
func (g *GridFileMock) Write(data []byte) (int, error) {
	return g.WriteFunc(data)
}

/*
IterMock is a mock Iter
*/
type IterMock struct {
	AllFunc     func(result interface{}) error
	CloseFunc   func() error
	DoneFunc    func() bool
	ErrFunc     func() error
	ForFunc     func(result interface{}, f func() error) error
	NextFunc    func(result interface{}) bool
	StateFunc   func() (int64, []bson.Raw)
	TimeoutFunc func() bool
}

/*
All is a mock function
*/
func (i *IterMock) All(result interface{}) error {
	return i.AllFunc(result)
}

/*
All is a mock function
*/
func (i *IterMock) Close() error {
	return i.CloseFunc()
}

/*
All is a mock function
*/
func (i *IterMock) Done() bool {
	return i.DoneFunc()
}

/*
All is a mock function
*/
func (i *IterMock) Err() error {
	return i.ErrFunc()
}

/*
All is a mock function
*/
func (i *IterMock) For(result interface{}, f func() error) error {
	return i.ForFunc(result, f)
}

/*
All is a mock function
*/
func (i *IterMock) Next(result interface{}) bool {
	return i.NextFunc(result)
}

/*
All is a mock function
*/
func (i *IterMock) State() (int64, []bson.Raw) {
	return i.StateFunc()
}

/*
All is a mock function
*/
func (i *IterMock) Timeout() bool {
	return i.TimeoutFunc()
}

/*
WriteToResultInterface writes a value to an interface
*/
func WriteToResultInterface(value, result interface{}) error {
	var err error
	var b bytes.Buffer

	encoder := gob.NewEncoder(&b)
	decoder := gob.NewDecoder(&b)

	if err = encoder.Encode(value); err != nil {
		return fmt.Errorf("Error encoding value in WriteToResultInterface: %w", err)
	}

	if err = decoder.Decode(result); err != nil {
		return fmt.Errorf("Error decoding value to result in WriteToResultInterface: %w", err)
	}

	return nil
}
