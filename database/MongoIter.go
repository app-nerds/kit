package database

type MongoIter interface {
	All(result interface{}) error
	Close() error
	Err() error
	For(result interface{}, f func() error) error
	Next(result interface{}) bool
	Timeout() bool
}
