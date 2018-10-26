package database

type MongoSession interface {
	Close()
	DB(name string) MongoDatabase
}
