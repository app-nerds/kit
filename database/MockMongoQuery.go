package database

/*
MockMongoQuery is used for unit testing MongoDB queries
*/
type MockMongoQuery struct {
	AllFunc      func(result interface{}) error
	CountFunc    func() (int, error)
	DistinctFunc func(key string, result interface{}) error
	LimitFunc    func(n int) MongoQuery
	OneFunc      func(result interface{}) error
	SelectFunc   func(selector interface{}) MongoQuery
	SkipFunc     func(n int) MongoQuery
	SortFunc     func(fields ...string) MongoQuery
}

func (m *MockMongoQuery) All(result interface{}) error {
	return m.AllFunc(result)
}

func (m *MockMongoQuery) Count() (int, error) {
	return m.CountFunc()
}

func (m *MockMongoQuery) Distinct(key string, result interface{}) error {
	return m.DistinctFunc(key, result)
}

func (m *MockMongoQuery) Limit(n int) MongoQuery {
	return m.LimitFunc(n)
}

func (m *MockMongoQuery) One(result interface{}) error {
	return m.OneFunc(result)
}

func (m *MockMongoQuery) Select(selector interface{}) MongoQuery {
	return m.SelectFunc(selector)
}

func (m *MockMongoQuery) Skip(n int) MongoQuery {
	return m.SkipFunc(n)
}

func (m *MockMongoQuery) Sort(fields ...string) MongoQuery {
	return m.SortFunc(fields...)
}
