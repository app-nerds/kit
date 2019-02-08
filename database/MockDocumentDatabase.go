package database

import "github.com/globalsign/mgo"

type MockDocumentDatabase struct {
	ConnectFunc    func(connection *Connection) error
	DisconnectFunc func()

	GetCollectionFunc func(name string) *mgo.Collection
	GetDBFunc         func() *mgo.Database
	GridFSFunc        func(prefix string) *mgo.GridFS
}

func (m *MockDocumentDatabase) Connect(connection *Connection) error {
	return m.ConnectFunc(connection)
}

func (m *MockDocumentDatabase) Disconnect() {
	m.DisconnectFunc()
}

func (m *MockDocumentDatabase) GetCollection(name string) *mgo.Collection {
	return m.GetCollectionFunc(name)
}

func (m *MockDocumentDatabase) GetDB() *mgo.Database {
	return m.GetDBFunc()
}

func (m *MockDocumentDatabase) GridFS(prefix string) *mgo.GridFS {
	return m.GridFSFunc(prefix)
}
