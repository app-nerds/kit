package database

import "github.com/globalsign/mgo"

type MongoDatabase interface {
	C(name string) MongoCollection
	GridFS(prefix string) *mgo.GridFS
}
