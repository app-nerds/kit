package database

import (
	"github.com/globalsign/mgo"
)

/*
DocumentDatabase defines the basics of what a document-based database
system can do
*/
type DocumentDatabase interface {
	BaseDatabase

	GetCollection(name string) *mgo.Collection
	GetDB() *mgo.Database
	GridFS(prefix string) *mgo.GridFS
}
