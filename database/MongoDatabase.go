package database

type MongoDatabase interface {
	C(name string) MongoCollection
	GridFS(prefix string) MongoGridFS
}
