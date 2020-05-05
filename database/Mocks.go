/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package database

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

type SessionMock struct {
	DBFunc    func(name string) Database
	CloseFunc func()
}

func (s *SessionMock) DB(name string) Database {
	return s.DBFunc(name)
}

func (s *SessionMock) Close() {
	s.CloseFunc()
}

type DatabaseMock struct {
	CFunc      func(name string) Collection
	GridFSFunc func(prefix string) GridFS
}

func (d *DatabaseMock) C(name string) Collection {
	return d.CFunc(name)
}

func (d *DatabaseMock) GridFS(prefix string) GridFS {
	return d.GridFSFunc(prefix)
}

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

func (c *CollectionMock) Count() (int, error) {
	return c.CountFunc()
}

func (c *CollectionMock) DropAllIndexes() error {
	return c.DropAllIndexesFunc()
}

func (c *CollectionMock) DropCollection() error {
	return c.DropCollectionFunc()
}

func (c *CollectionMock) DropIndex(key ...string) error {
	return c.DropIndexFunc(key...)
}

func (c *CollectionMock) DropIndexName(name string) error {
	return c.DropIndexNameFunc(name)
}

func (c *CollectionMock) EnsureIndex(index mgo.Index) error {
	return c.EnsureIndexFunc(index)
}

func (c *CollectionMock) EnsureIndexKey(key ...string) error {
	return c.EnsureIndexKeyFunc(key...)
}

func (c *CollectionMock) Find(query interface{}) Query {
	return c.FindFunc(query)
}

func (c *CollectionMock) FindId(id interface{}) Query {
	return c.FindIdFunc(id)
}

func (c *CollectionMock) Indexes() ([]mgo.Index, error) {
	return c.IndexesFunc()
}

func (c *CollectionMock) Insert(docs ...interface{}) error {
	return c.InsertFunc(docs...)
}

func (c *CollectionMock) Remove(selector interface{}) error {
	return c.RemoveFunc(selector)
}

func (c *CollectionMock) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	return c.RemoveAllFunc(selector)
}

func (c *CollectionMock) RemoveId(id interface{}) error {
	return c.RemoveIdFunc(id)
}

func (c *CollectionMock) Update(selector interface{}, update interface{}) error {
	return c.UpdateFunc(selector, update)
}

func (c *CollectionMock) UpdateAll(selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpdateAllFunc(selector, update)
}

func (c *CollectionMock) UpdateId(id interface{}, update interface{}) error {
	return c.UpdateIdFunc(id, update)
}

func (c *CollectionMock) Upsert(selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertFunc(selector, update)
}

func (c *CollectionMock) UpsertId(id interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertIdFunc(id, update)
}

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

func (q *QueryMock) All(result interface{}) error {
	return q.AllFunc(result)
}

func (q *QueryMock) Count() (int, error) {
	return q.CountFunc()
}

func (q *QueryMock) Distinct(key string, result interface{}) error {
	return q.DistinctFunc(key, result)
}

func (q *QueryMock) Limit(n int) Query {
	return q.LimitFunc(n)
}

func (q *QueryMock) One(result interface{}) error {
	return q.OneFunc(result)
}

func (q *QueryMock) Select(selector interface{}) Query {
	return q.SelectFunc(selector)
}

func (q *QueryMock) Skip(n int) Query {
	return q.SkipFunc(n)
}

func (q *QueryMock) Sort(fields ...string) Query {
	return q.SortFunc(fields...)
}

type GridFSMock struct {
	CreateFunc   func(name string) (GridFile, error)
	FindFunc     func(query interface{}) Query
	OpenFunc     func(name string) (GridFile, error)
	OpenIdFunc   func(id interface{}) (GridFile, error)
	RemoveFunc   func(name string) (err error)
	RemoveIdFunc func(id interface{}) error
}

func (g *GridFSMock) Create(name string) (GridFile, error) {
	return g.CreateFunc(name)
}

func (g *GridFSMock) Find(query interface{}) Query {
	return g.FindFunc(query)
}

func (g *GridFSMock) Open(name string) (GridFile, error) {
	return g.OpenFunc(name)
}

func (g *GridFSMock) OpenId(id interface{}) (GridFile, error) {
	return g.OpenIdFunc(id)
}

func (g *GridFSMock) Remove(name string) (err error) {
	return g.RemoveFunc(name)
}

func (g *GridFSMock) RemoveId(id interface{}) error {
	return g.RemoveIdFunc(id)
}

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

func (g *GridFileMock) Abort() {
	g.AbortFunc()
}

func (g *GridFileMock) Close() error {
	return g.CloseFunc()
}

func (g *GridFileMock) ContentType() string {
	return g.ContentTypeFunc()
}

func (g *GridFileMock) GetMeta(result interface{}) error {
	return g.GetMetaFunc(result)
}

func (g *GridFileMock) Id() interface{} {
	return g.IdFunc()
}

func (g *GridFileMock) MD5() string {
	return g.MD5Func()
}

func (g *GridFileMock) Name() string {
	return g.NameFunc()
}

func (g *GridFileMock) Read(b []byte) (int, error) {
	return g.ReadFunc(b)
}

func (g *GridFileMock) Seek(offset int64, whence int) (int64, error) {
	return g.SeekFunc(offset, whence)
}

func (g *GridFileMock) SetChunkSize(bytes int) {
	g.SetChunkSizeFunc(bytes)
}

func (g *GridFileMock) SetContentType(ctype string) {
	g.SetContentTypeFunc(ctype)
}

func (g *GridFileMock) SetId(id interface{}) {
	g.SetIdFunc(id)
}

func (g *GridFileMock) SetMeta(metadata interface{}) {
	g.SetMetaFunc(metadata)
}

func (g *GridFileMock) SetName(name string) {
	g.SetName(name)
}

func (g *GridFileMock) SetUploadDate(t time.Time) {
	g.SetUploadDateFunc(t)
}

func (g *GridFileMock) Size() int64 {
	return g.SizeFunc()
}

func (g *GridFileMock) Write(data []byte) (int, error) {
	return g.WriteFunc(data)
}

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

func (i *IterMock) All(result interface{}) error {
	return i.AllFunc(result)
}

func (i *IterMock) Close() error {
	return i.CloseFunc()
}

func (i *IterMock) Done() bool {
	return i.DoneFunc()
}

func (i *IterMock) Err() error {
	return i.ErrFunc()
}

func (i *IterMock) For(result interface{}, f func() error) error {
	return i.ForFunc(result, f)
}

func (i *IterMock) Next(result interface{}) bool {
	return i.NextFunc(result)
}

func (i *IterMock) State() (int64, []bson.Raw) {
	return i.StateFunc()
}

func (i *IterMock) Timeout() bool {
	return i.TimeoutFunc()
}

func WriteToResultInterface(value, result interface{}) error {
	var err error
	var b bytes.Buffer

	encoder := gob.NewEncoder(&b)
	decoder := gob.NewDecoder(&b)

	if err = encoder.Encode(value); err != nil {
		return errors.Wrapf(err, "Error encoding value in WriteToResultInterface")
	}

	if err = decoder.Decode(result); err != nil {
		return errors.Wrapf(err, "Error decoding value to result in WriteToResultInterface")
	}

	return nil
}
