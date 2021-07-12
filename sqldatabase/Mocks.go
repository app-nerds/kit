package sqldatabase

import (
	"context"
	"database/sql"
	"reflect"
	"time"
)

type MockColumnType struct {
	DatabaseTypeNameFunc func() string
	DecimalSizeFunc      func() (precision, scale int64, ok bool)
	LengthFunc           func() (length int64, ok bool)
	NameFunc             func() string
	NullableFunc         func() (nullable, ok bool)
	ScanTypeFunc         func() reflect.Type
}

type MockDB struct {
	BeginFunc              func() (Tx, error)
	CloseFunc              func() error
	ExecFunc               func(query string, args ...interface{}) (sql.Result, error)
	ExecContextFunc        func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PingFunc               func() error
	PingContextFunc        func(ctx context.Context) error
	PrepareFunc            func(query string) (Stmt, error)
	PrepareContextFunc     func(ctx context.Context, query string) (Stmt, error)
	QueryFunc              func(query string, args ...interface{}) (Rows, error)
	QueryContextFunc       func(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRowFunc           func(query string, args ...interface{}) Row
	QueryRowContextFunc    func(ctx context.Context, query string, args ...interface{}) Row
	SetConnMaxIdleTimeFunc func(d time.Duration)
	SetConnMaxLifetimeFunc func(d time.Duration)
	SetMaxIdleConnsFunc    func(n int)
	SetMaxOpenConnsFunc    func(n int)
}

type MockRow struct {
	ErrFunc  func() error
	ScanFunc func(dest ...interface{}) error
}

type MockRows struct {
	CloseFunc         func() error
	ColumnTypesFunc   func() ([]ColumnType, error)
	ColumnsFunc       func() ([]string, error)
	ErrFunc           func() error
	NextFunc          func() bool
	NextResultSetFunc func() bool
	ScanFunc          func(dst ...interface{}) error
}

type MockStmt struct {
	CloseFunc           func() error
	ExecFunc            func(args ...interface{}) (sql.Result, error)
	ExecContextFunc     func(ctx context.Context, args ...interface{}) (sql.Result, error)
	QueryFunc           func(args ...interface{}) (Rows, error)
	QueryContextFunc    func(ctx context.Context, args ...interface{}) (Rows, error)
	QueryRowFunc        func(args ...interface{}) Row
	QueryRowContextFunc func(ctx context.Context, args ...interface{}) Row
}

type MockTx struct {
	CommitFunc          func() error
	ExecFunc            func(query string, args ...interface{}) (sql.Result, error)
	ExecContextFunc     func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareFunc         func(query string) (Stmt, error)
	PrepareContextFunc  func(ctx context.Context, query string) (Stmt, error)
	QueryFunc           func(query string, args ...interface{}) (Rows, error)
	QueryContextFunc    func(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRowFunc        func(query string, args ...interface{}) Row
	QueryRowContextFunc func(ctx context.Context, query string, args ...interface{}) Row
	RollbackFunc        func() error
	StmtFunc            func(stmt *sql.Stmt) Stmt
	StmtContextFunc     func(ctx context.Context, stmt *sql.Stmt) Stmt
}

/********************************************************************
 * MockColumnType
 *******************************************************************/
func (m *MockColumnType) DatabaseTypeName() string {
	return m.DatabaseTypeNameFunc()
}

func (m *MockColumnType) DecimalSize() (precision, scale int64, ok bool) {
	return m.DecimalSizeFunc()
}

func (m *MockColumnType) Length() (length int64, ok bool) {
	return m.LengthFunc()
}

func (m *MockColumnType) Name() string {
	return m.NameFunc()
}

func (m *MockColumnType) Nullable() (nullable, ok bool) {
	return m.NullableFunc()
}

func (m *MockColumnType) ScanType() reflect.Type {
	return m.ScanTypeFunc()
}

/********************************************************************
 * MockDB
 *******************************************************************/
func (m *MockDB) Begin() (Tx, error) {
	return m.BeginFunc()
}

func (m *MockDB) Close() error {
	return m.CloseFunc()
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.ExecFunc(query, args...)
}

func (m *MockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.ExecContextFunc(ctx, query, args...)
}

func (m *MockDB) Ping() error {
	return m.PingFunc()
}

func (m *MockDB) PingContext(ctx context.Context) error {
	return m.PingContextFunc(ctx)
}

func (m *MockDB) Prepare(query string) (Stmt, error) {
	return m.PrepareFunc(query)
}

func (m *MockDB) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	return m.PrepareContextFunc(ctx, query)
}

func (m *MockDB) Query(query string, args ...interface{}) (Rows, error) {
	return m.QueryFunc(query, args...)
}

func (m *MockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	return m.QueryContextFunc(ctx, query, args...)
}

func (m *MockDB) QueryRow(query string, args ...interface{}) Row {
	return m.QueryRowFunc(query, args...)
}

func (m *MockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return m.QueryRowContextFunc(ctx, query, args...)
}

func (m *MockDB) SetConnMaxIdleTime(d time.Duration) {
	m.SetConnMaxIdleTimeFunc(d)
}

func (m *MockDB) SetConnMaxLifetime(d time.Duration) {
	m.SetConnMaxLifetimeFunc(d)
}

func (m *MockDB) SetMaxIdleConns(n int) {
	m.SetMaxIdleConnsFunc(n)
}

func (m *MockDB) SetMaxOpenConns(n int) {
	m.SetMaxOpenConnsFunc(n)
}

/********************************************************************
 * MockRow
 *******************************************************************/
func (m *MockRow) Err() error {
	return m.ErrFunc()
}

func (m *MockRow) Scan(dest ...interface{}) error {
	return m.ScanFunc(dest...)
}

/********************************************************************
 * MockRows
 *******************************************************************/
func (m *MockRows) Close() error {
	return m.CloseFunc()
}

func (m *MockRows) ColumnTypes() ([]ColumnType, error) {
	return m.ColumnTypesFunc()
}

func (m *MockRows) Columns() ([]string, error) {
	return m.ColumnsFunc()
}

func (m *MockRows) Err() error {
	return m.ErrFunc()
}

func (m *MockRows) Next() bool {
	return m.NextFunc()
}

func (m *MockRows) NextResultSet() bool {
	return m.NextResultSetFunc()
}

func (m *MockRows) Scan(dst ...interface{}) error {
	return m.ScanFunc(dst...)
}

/********************************************************************
 * MockStmt
 *******************************************************************/
func (m *MockStmt) Close() error {
	return m.CloseFunc()
}

func (m *MockStmt) Exec(args ...interface{}) (sql.Result, error) {
	return m.ExecFunc(args...)
}

func (m *MockStmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return m.ExecContextFunc(ctx, args...)
}

func (m *MockStmt) Query(args ...interface{}) (Rows, error) {
	return m.QueryFunc(args...)
}

func (m *MockStmt) QueryContext(ctx context.Context, args ...interface{}) (Rows, error) {
	return m.QueryContextFunc(ctx, args...)
}

func (m *MockStmt) QueryRow(args ...interface{}) Row {
	return m.QueryRowFunc(args...)
}

func (m *MockStmt) QueryRowContext(ctx context.Context, args ...interface{}) Row {
	return m.QueryRowContextFunc(ctx, args...)
}

/********************************************************************
 * MockTx
 *******************************************************************/
func (m *MockTx) Commit() error {
	return m.CommitFunc()
}

func (m *MockTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.ExecFunc(query, args...)
}

func (m *MockTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.ExecContextFunc(ctx, query, args...)
}

func (m *MockTx) Prepare(query string) (Stmt, error) {
	return m.PrepareFunc(query)
}

func (m *MockTx) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	return m.PrepareContextFunc(ctx, query)
}

func (m *MockTx) Query(query string, args ...interface{}) (Rows, error) {
	return m.QueryFunc(query, args...)
}

func (m *MockTx) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	return m.QueryContextFunc(ctx, query, args...)
}

func (m *MockTx) QueryRow(query string, args ...interface{}) Row {
	return m.QueryRowFunc(query, args...)
}

func (m *MockTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	return m.QueryRowContextFunc(ctx, query, args...)
}

func (m *MockTx) Rollback() error {
	return m.RollbackFunc()
}

func (m *MockTx) Stmt(stmt *sql.Stmt) Stmt {
	return m.StmtFunc(stmt)
}

func (m *MockTx) StmtContext(ctx context.Context, stmt *sql.Stmt) Stmt {
	return m.StmtContextFunc(ctx, stmt)
}
