package database

import (
	"github.com/globalsign/mgo"
)

/*
MongoCollection is an interface describing MongoDB collection methods
*/
type MongoCollection interface {
	Count() (int, error)
	EnsureIndex(index mgo.Index) error
	Find(query interface{}) *MongoQueryWrapper
	FindId(id interface{}) *mgo.Query
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (*mgo.ChangeInfo, error)
	RemoveId(id interface{}) error
	Update(selector, update interface{}) error
	UpdateAll(selector, update interface{}) (*mgo.ChangeInfo, error)
	UpdateId(id, update interface{}) error
	Upsert(selector, update interface{}) (*mgo.ChangeInfo, error)
	UpsertId(id, update interface{}) (*mgo.ChangeInfo, error)
}
