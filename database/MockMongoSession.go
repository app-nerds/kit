package database

type MockMongoSession struct {
	DBFunc func(name string) MongoDatabase
}

func (m *MockMongoSession) Close() {
}

func (m *MockMongoSession) DB(name string) MongoDatabase {
	return m.DBFunc(name)
}
