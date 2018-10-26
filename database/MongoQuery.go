package database

import (
	"github.com/globalsign/mgo"
)

/*
MongoQuery is an interface describing a MongoDB query
*/
type MongoQuery interface {
	All(result interface{}) error
	Count() (int, error)
	Distinct(key string, result interface{}) error
	Limit(n int) MongoQuery
	One(result interface{}) error
	Select(selector interface{}) MongoQuery
	Skip(n int) MongoQuery
	Sort(fields ...string) MongoQuery
}

type MongoQueryImpl struct {
	*mgo.Query
}
