package database

type MongoGridFS interface {
	Create(name string) (MongoGridFile, error)
	Find(query interface{}) MongoQuery
	Open(name string) (MongoGridFile, error)
	OpenId(id interface{}) (MongoGridFile, error)
	OpenNext(MongoIter, MongoGridFile) bool
	Remove(name string) error
	RemoveId(id interface{}) error
}
