package database

type MockMongoDatabase struct {
	CFunc      func(name string) MongoCollection
	GridFSFunc func(prefix string) MongoGridFS
}

func (m *MockMongoDatabase) C(name string) MongoCollection {
	return m.CFunc(name)
}

func (m *MockMongoDatabase) GridFS(prefix string) MongoGridFS {
	return m.GridFSFunc(prefix)
}
