package database

import "github.com/globalsign/mgo"

/*
MockMongoCollection is used for unit testing to mock a MongoDB
*/
type MockMongoCollection struct {
	CountFunc       func() (int, error)
	EnsureIndexFunc func(index mgo.Index) error
	FindFunc        func(query interface{}) *mgo.Query
	FindIdFunc      func(id interface{}) *mgo.Query
	InsertFunc      func(docs ...interface{}) error
	RemoveFunc      func(selector interface{}) error
	RemoveAllFunc   func(selector interface{}) (*mgo.ChangeInfo, error)
	RemoveIdFunc    func(id interface{}) error
	UpdateFunc      func(selector, update interface{}) error
	UpdateAllFunc   func(selector, update interface{}) (*mgo.ChangeInfo, error)
	UpdateIdFunc    func(id, update interface{}) error
	UpsertFunc      func(selector, update interface{}) (*mgo.ChangeInfo, error)
	UpsertIdFunc    func(id, update interface{}) (*mgo.ChangeInfo, error)
}

func (c *MockMongoCollection) Count() (int, error) {
	return c.CountFunc()
}

func (c *MockMongoCollection) EnsureIndex(index mgo.Index) error {
	return c.EnsureIndexFunc(index)
}

func (c *MockMongoCollection) Find(query interface{}) *mgo.Query {
	return c.FindFunc(query)
}

func (c *MockMongoCollection) FindId(id interface{}) *mgo.Query {
	return c.FindIdFunc(id)
}

func (c *MockMongoCollection) Insert(docs ...interface{}) error {
	return c.InsertFunc(docs)
}

func (c *MockMongoCollection) Remove(selector interface{}) error {
	return c.RemoveFunc(selector)
}

func (c *MockMongoCollection) RemoveId(id interface{}) error {
	return c.RemoveIdFunc(id)
}

func (c *MockMongoCollection) RemoveAll(selector interface{}) (*mgo.ChangeInfo, error) {
	return c.RemoveAllFunc(selector)
}

func (c *MockMongoCollection) Update(selector, update interface{}) error {
	return c.UpdateFunc(selector, update)
}

func (c *MockMongoCollection) UpdateAll(selector, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpdateAllFunc(selector, update)
}

func (c *MockMongoCollection) UpdateId(id, update interface{}) error {
	return c.UpdateIdFunc(id, update)
}

func (c *MockMongoCollection) Upsert(selector, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertFunc(selector, update)
}

func (c *MockMongoCollection) UpsertId(id, update interface{}) (*mgo.ChangeInfo, error) {
	return c.UpsertIdFunc(id, update)
}
