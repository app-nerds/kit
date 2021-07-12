package sqldatabase

import (
	"context"
	"database/sql"
	"reflect"
	"time"
)

/*
ColumnType contains the name and type of a column.
*/
type ColumnType interface {
	/*
		DatabaseTypeName returns the database system name of the column
		type. If an empty string is returned, then the driver type name
		is not supported. Consult your driver documentation for a list of
		driver data types. Length specifiers are not included. Common type
		names include "VARCHAR", "TEXT", "NVARCHAR", "DECIMAL", "BOOL",
		"INT", and "BIGINT".
	*/
	DatabaseTypeName() string

	/*
		DecimalSize returns the scale and precision of a decimal type.
		If not applicable or if not supported ok is false.
	*/
	DecimalSize() (precision, scale int64, ok bool)

	/*
		Length returns the column type length for variable length column types
		such as text and binary field types. If the type length is unbounded
		the value will be math.MaxInt64 (any database limits will still apply).
		If the column type is not variable length, such as an int, or if not
		supported by the driver ok is false.
	*/
	Length() (length int64, ok bool)

	/*
		Name returns the name or alias of the column.
	*/
	Name() string

	/*
		Nullable reports whether the column may be null. If a driver does
		not support this property ok will be false.
	*/
	Nullable() (nullable, ok bool)

	/*
		ScanType returns a Go type suitable for scanning into using Rows.Scan.
		If a driver does not support this property ScanType will return the
		type of an empty interface.
	*/
	ScanType() reflect.Type
}

/*
DB is a database handle representing a pool of zero or more underlying
connections. It's safe for concurrent use by multiple goroutines.

The sql package creates and frees connections automatically; it also
maintains a free pool of idle connections. If the database has a
concept of per-connection state, such state can be reliably observed
within a transaction (Tx) or connection (Conn). Once DB.Begin is called,
the returned Tx is bound to a single connection. Once Commit or Rollback
is called on the transaction, that transaction's connection is returned
to DB's idle connection pool. The pool size can be controlled with
SetMaxIdleConns.
*/
type DB interface {
	/*
		Begin starts a transaction. The default isolation level is dependent
		on the driver.
	*/
	Begin() (Tx, error)

	/*
		Close closes the database and prevents new queries from starting.
		Close then waits for all queries that have started processing on
		the server to finish.

		It is rare to Close a DB, as the DB handle is meant to be long-lived
		and shared between many goroutines.
	*/
	Close() error

	/*
		Exec executes a query without returning any rows. The args are for
		any placeholder parameters in the query.
	*/
	Exec(query string, args ...interface{}) (sql.Result, error)

	/*
		ExecContext executes a query without returning any rows. The args
		are for any placeholder parameters in the query.
	*/
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	/*
		Ping verifies a connection to the database is still alive,
		establishing a connection if necessary.
	*/
	Ping() error

	/*
		PingContext verifies a connection to the database is still alive,
		establishing a connection if necessary.
	*/
	PingContext(ctx context.Context) error

	/*
		Prepare creates a prepared statement for later queries or executions.
		Multiple queries or executions may be run concurrently from the
		returned statement. The caller must call the statement's Close method
		when the statement is no longer needed.
	*/
	Prepare(query string) (Stmt, error)

	/*
		PrepareContext creates a prepared statement for later queries or executions.
		Multiple queries or executions may be run concurrently from the returned
		statement. The caller must call the statement's Close method when the
		statement is no longer needed.

		The provided context is used for the preparation of the statement, not
		for the execution of the statement.
	*/
	PrepareContext(ctx context.Context, query string) (Stmt, error)

	/*
		Query executes a query that returns rows, typically a SELECT. The args are
		for any placeholder parameters in the query.
	*/
	Query(query string, args ...interface{}) (Rows, error)

	/*
		QueryContext executes a query that returns rows, typically a SELECT.
		The args are for any placeholder parameters in the query.
	*/
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)

	/*
		QueryRow executes a query that is expected to return at most one row.
		QueryRow always returns a non-nil value. Errors are deferred until
		Row's Scan method is called. If the query selects no rows, the *Row's
		Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first
		selected row and discards the rest.
	*/
	QueryRow(query string, args ...interface{}) Row

	/*
		QueryRowContext executes a query that is expected to return at most
		one row. QueryRowContext always returns a non-nil value. Errors are
		deferred until Row's Scan method is called. If the query selects no
		rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's
		Scan scans the first selected row and discards the rest.
	*/
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row

	/*
		SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.

		Expired connections may be closed lazily before reuse.

		If d <= 0, connections are not closed due to a connection's idle time.
	*/
	SetConnMaxIdleTime(d time.Duration)

	/*
		SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

		Expired connections may be closed lazily before reuse.

		If d <= 0, connections are not closed due to a connection's age.
	*/
	SetConnMaxLifetime(d time.Duration)

	/*
		SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

		If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new
		MaxIdleConns will be reduced to match the MaxOpenConns limit.

		If n <= 0, no idle connections are retained.

		The default max idle connections is currently 2. This may change in a future release.
	*/
	SetMaxIdleConns(n int)

	/*
		SetMaxOpenConns sets the maximum number of open connections to the database.

		If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
		MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.

		If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
	*/
	SetMaxOpenConns(n int)
}

/*
Row is the result of calling QueryRow to select a single row.
*/
type Row interface {
	/*
		Err provides a way for wrapping packages to check for query errors
		without calling Scan. Err returns the error, if any, that was
		encountered while running the query. If this error is not nil, this
		error will also be returned from Scan.
	*/
	Err() error

	/*
		Scan copies the columns from the matched row into the values pointed
		at by dest. See the documentation on Rows.Scan for details. If more
		than one row matches the query, Scan uses the first row and discards
		the rest. If no row matches the query, Scan returns ErrNoRows.
	*/
	Scan(dest ...interface{}) error
}

/*
Rows is the result of a query. Its cursor starts before the first row of
the result set. Use Next to advance from row to row.
*/
type Rows interface {
	/*
		Close closes the Rows, preventing further enumeration. If Next is
		called and returns false and there are no further result sets, the
		Rows are closed automatically and it will suffice to check the
		result of Err. Close is idempotent and does not affect the result
		of Err.
	*/
	Close() error

	/*
		ColumnTypes returns column information such as column type, length,
		and nullable. Some information may not be available from some drivers.
	*/
	ColumnTypes() ([]ColumnType, error)

	/*
		Columns returns the column names. Columns returns an error if the
		rows are closed.
	*/
	Columns() ([]string, error)

	/*
		Err returns the error, if any, that was encountered during iteration.
		Err may be called after an explicit or implicit Close.
	*/
	Err() error

	/*
		Next prepares the next result row for reading with the Scan method. It
		returns true on success, or false if there is no next result row or an
		error happened while preparing it. Err should be consulted to
		distinguish between the two cases.

		Every call to Scan, even the first one, must be preceded by a call to Next.
	*/
	Next() bool

	/*
		NextResultSet prepares the next result set for reading. It reports
		whether there is further result sets, or false if there is no further
		result set or if there is an error advancing to it. The Err method
		should be consulted to distinguish between the two cases.

		After calling NextResultSet, the Next method should always be called
		before scanning. If there are further result sets they may not have
		rows in the result set.
	*/
	NextResultSet() bool

	/*
		Scan copies the columns in the current row into the values pointed at by
		dest. The number of values in dest must be the same as the number of
		columns in Rows.
	*/
	Scan(dst ...interface{}) error
}

/*
Scanner describes a strut that can scan columns from a record
*/
type Scanner interface {
	/*
		Scan copies the columns into values pointed at by dest. The number
		of values in dest must be the same as the number of
		columns.
	*/
	Scan(dst ...interface{}) error
}

/*
Stmt is a prepared statement. A Stmt is safe for concurrent use by multiple
goroutines.

If a Stmt is prepared on a Tx or Conn, it will be bound to a single
underlying connection forever. If the Tx or Conn closes, the Stmt will
become unusable and all operations will return an error. If a Stmt is
prepared on a DB, it will remain usable for the lifetime of the DB. When
the Stmt needs to execute on a new underlying connection, it will prepare
itself on the new connection automatically.
*/
type Stmt interface {
	/*
		Close closes the statement.
	*/
	Close() error

	/*
		Exec executes a prepared statement with the given arguments and
		returns a Result summarizing the effect of the statement.
	*/
	Exec(args ...interface{}) (sql.Result, error)

	/*
		ExecContext executes a prepared statement with the given arguments
		and returns a Result summarizing the effect of the statement.
	*/
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)

	/*
		Query executes a prepared query statement with the given arguments
		and returns the query results as a *Rows.
	*/
	Query(args ...interface{}) (Rows, error)

	/*
		QueryContext executes a prepared query statement with the given
		arguments and returns the query results as a *Rows.
	*/
	QueryContext(ctx context.Context, args ...interface{}) (Rows, error)

	/*
		QueryRow executes a prepared query statement with the given arguments.
		If an error occurs during the execution of the statement, that error
		will be returned by a call to Scan on the returned *Row, which is always
		non-nil. If the query selects no rows, the *Row's Scan will return
		ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and
		discards the rest.
	*/
	QueryRow(args ...interface{}) Row

	/*
		QueryRowContext executes a prepared query statement with the given
		arguments. If an error occurs during the execution of the statement,
		that error will be returned by a call to Scan on the returned *Row,
		which is always non-nil. If the query selects no rows, the *Row's
		Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the
		first selected row and discards the rest.
	*/
	QueryRowContext(ctx context.Context, args ...interface{}) Row
}

/*
Tx is an in-progress database transaction.

A transaction must end with a call to Commit or Rollback.

After a call to Commit or Rollback, all operations on the transaction
fail with ErrTxDone.

The statements prepared for a transaction by calling the transaction's
Prepare or Stmt methods are closed by the call to Commit or Rollback.
*/
type Tx interface {
	/*
		Commit commits the transaction.
	*/
	Commit() error

	/*
		Exec executes a query that doesn't return rows. For example: an INSERT
		and UPDATE.
	*/
	Exec(query string, args ...interface{}) (sql.Result, error)

	/*
		ExecContext executes a query that doesn't return rows. For example: an
		INSERT and UPDATE.
	*/
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	/*
		Prepare creates a prepared statement for use within a transaction.

		The returned statement operates within the transaction and can no
		longer be used once the transaction has been committed or rolled back.

		To use an existing prepared statement on this transaction, see Tx.Stmt.
	*/
	Prepare(query string) (Stmt, error)

	/*
		PrepareContext creates a prepared statement for use within a
		transaction.

		The returned statement operates within the transaction and will
		be closed when the transaction has been committed or rolled back.

		To use an existing prepared statement on this transaction, see Tx.Stmt.

		The provided context will be used for the preparation of the
		context, not for the execution of the returned statement. The
		returned statement will run in the transaction context.
	*/
	PrepareContext(ctx context.Context, query string) (Stmt, error)

	/*
		Query executes a query that returns rows, typically a SELECT.
	*/
	Query(query string, args ...interface{}) (Rows, error)

	/*
		QueryContext executes a query that returns rows, typically a SELECT.
	*/
	QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)

	/*
		QueryRow executes a query that is expected to return at most one row.
		QueryRow always returns a non-nil value. Errors are deferred until
		Row's Scan method is called. If the query selects no rows, the *Row's
		Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first
		selected row and discards the rest.
	*/
	QueryRow(query string, args ...interface{}) Row

	/*
		QueryRowContext executes a query that is expected to return at most
		one row. QueryRowContext always returns a non-nil value. Errors are
		deferred until Row's Scan method is called. If the query selects no
		rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's
		Scan scans the first selected row and discards the rest.
	*/
	QueryRowContext(ctx context.Context, query string, args ...interface{}) Row

	/*
		Rollback aborts the transaction.
	*/
	Rollback() error

	/*
		Stmt returns a transaction-specific prepared statement from an existing statement.
	*/
	Stmt(stmt *sql.Stmt) Stmt

	/*
		StmtContext returns a transaction-specific prepared statement from an
		existing statement.
	*/
	StmtContext(ctx context.Context, stmt *sql.Stmt) Stmt
}

type sqlColumnType struct {
	*sql.ColumnType
}

type sqlDB struct {
	*sql.DB
}

type sqlRow struct {
	*sql.Row
}

type sqlRows struct {
	*sql.Rows
}

type sqlStmt struct {
	*sql.Stmt
}

type sqlTx struct {
	*sql.Tx
}

/*
Open opens a database specified by its database driver name and a driver-specific
data source name, usually consisting of at least a database name and connection
information.

Most users will open a database via a driver-specific connection helper function
that returns a *DB. No database drivers are included in the Go standard library.
See https://golang.org/s/sqldrivers for a list of third-party drivers.

Open may just validate its arguments without creating a connection to the
database. To verify that the data source name is valid, call Ping.

The returned DB is safe for concurrent use by multiple goroutines and
maintains its own pool of idle connections. Thus, the Open function should be
called just once. It is rarely necessary to close a DB.
*/
func Open(driverName, dataSourceName string) (DB, error) {
	db, err := sql.Open(driverName, dataSourceName)

	return &sqlDB{
		DB: db,
	}, err
}

/********************************************************************
 * sqlColumnType
 *******************************************************************/
func (s *sqlColumnType) DatabaseTypeName() string {
	return s.ColumnType.DatabaseTypeName()
}

func (s *sqlColumnType) DecimalSize() (precision, scale int64, ok bool) {
	return s.ColumnType.DecimalSize()
}

func (s *sqlColumnType) Length() (length int64, ok bool) {
	return s.ColumnType.Length()
}

func (s *sqlColumnType) Name() string {
	return s.ColumnType.Name()
}

func (s *sqlColumnType) Nullable() (nullable, ok bool) {
	return s.ColumnType.Nullable()
}

func (s *sqlColumnType) ScanType() reflect.Type {
	return s.ColumnType.ScanType()
}

/********************************************************************
 * sqlDB
 *******************************************************************/
func (s *sqlDB) Begin() (Tx, error) {
	tx, err := s.DB.Begin()
	return &sqlTx{
		Tx: tx,
	}, err
}

func (s *sqlDB) Close() error {
	return s.DB.Close()
}

func (s *sqlDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *sqlDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.DB.ExecContext(ctx, query, args...)
}

func (s *sqlDB) Ping() error {
	return s.DB.Ping()
}

func (s *sqlDB) PingContext(ctx context.Context) error {
	return s.DB.PingContext(ctx)
}

func (s *sqlDB) Prepare(query string) (Stmt, error) {
	stmt, err := s.DB.Prepare(query)
	return &sqlStmt{
		Stmt: stmt,
	}, err
}

func (s *sqlDB) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	stmt, err := s.DB.PrepareContext(ctx, query)
	return &sqlStmt{
		Stmt: stmt,
	}, err
}

func (s *sqlDB) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := s.DB.Query(query, args...)

	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlDB) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := s.DB.QueryContext(ctx, query, args...)
	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlDB) QueryRow(query string, args ...interface{}) Row {
	row := s.DB.QueryRow(query, args...)
	return &sqlRow{
		Row: row,
	}
}

func (s *sqlDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	row := s.DB.QueryRowContext(ctx, query, args...)
	return &sqlRow{
		Row: row,
	}
}

func (s *sqlDB) SetConnMaxIdleTime(d time.Duration) {
	s.DB.SetConnMaxIdleTime(d)
}

func (s *sqlDB) SetConnMaxLifetime(d time.Duration) {
	s.DB.SetConnMaxLifetime(d)
}

func (s *sqlDB) SetMaxIdleConns(n int) {
	s.DB.SetMaxIdleConns(n)
}

func (s *sqlDB) SetMaxOpenConns(n int) {
	s.DB.SetMaxOpenConns(n)
}

/********************************************************************
* sqlRows
*******************************************************************/
func (s *sqlRows) Close() error {
	return s.Rows.Close()
}

func (s *sqlRows) ColumnTypes() ([]ColumnType, error) {
	columnTypes, err := s.Rows.ColumnTypes()
	result := make([]ColumnType, 0, 10)

	for _, c := range columnTypes {
		result = append(result, &sqlColumnType{
			ColumnType: c,
		})
	}

	return result, err
}

func (s *sqlRows) Columns() ([]string, error) {
	return s.Rows.Columns()
}

func (s *sqlRows) Err() error {
	return s.Rows.Err()
}

func (s *sqlRows) Next() bool {
	return s.Rows.Next()
}

func (s *sqlRows) NextResultSet() bool {
	return s.Rows.NextResultSet()
}

func (s *sqlRows) Scan(dst ...interface{}) error {
	return s.Rows.Scan(dst...)
}

/********************************************************************
 * sqlStmt
 *******************************************************************/
func (s *sqlStmt) Close() error {
	return s.Stmt.Close()
}

func (s *sqlStmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.Stmt.Exec(args...)
}

func (s *sqlStmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.Stmt.ExecContext(ctx, args...)
}

func (s *sqlStmt) Query(args ...interface{}) (Rows, error) {
	rows, err := s.Stmt.Query(args...)
	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlStmt) QueryContext(ctx context.Context, args ...interface{}) (Rows, error) {
	rows, err := s.Stmt.QueryContext(ctx, args...)
	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlStmt) QueryRow(args ...interface{}) Row {
	row := s.Stmt.QueryRow(args...)
	return &sqlRow{
		Row: row,
	}
}

func (s *sqlStmt) QueryRowContext(ctx context.Context, args ...interface{}) Row {
	row := s.Stmt.QueryRowContext(ctx, args...)
	return &sqlRow{
		Row: row,
	}
}

/********************************************************************
 * sqlTx
 *******************************************************************/

/*
Commit commits the transaction.
*/
func (s *sqlTx) Commit() error {
	return s.Tx.Commit()
}

/*
Exec executes a query that doesn't return rows. For example: an INSERT
and UPDATE.
*/
func (s *sqlTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.Tx.Exec(query, args...)
}

/*
ExecContext executes a query that doesn't return rows. For example: an INSERT
and UPDATE.
*/
func (s *sqlTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.Tx.ExecContext(ctx, query, args...)
}

func (s *sqlTx) Prepare(query string) (Stmt, error) {
	stmt, err := s.Tx.Prepare(query)
	return &sqlStmt{
		Stmt: stmt,
	}, err
}

func (s *sqlTx) PrepareContext(ctx context.Context, query string) (Stmt, error) {
	stmt, err := s.Tx.PrepareContext(ctx, query)
	return &sqlStmt{
		Stmt: stmt,
	}, err
}

func (s *sqlTx) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := s.Tx.Query(query, args...)
	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlTx) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := s.Tx.QueryContext(ctx, query, args...)
	return &sqlRows{
		Rows: rows,
	}, err
}

func (s *sqlTx) QueryRow(query string, args ...interface{}) Row {
	row := s.Tx.QueryRow(query, args...)
	return &sqlRow{
		Row: row,
	}
}

func (s *sqlTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) Row {
	row := s.Tx.QueryRowContext(ctx, query, args...)
	return &sqlRow{
		Row: row,
	}
}

func (s *sqlTx) Rollback() error {
	return s.Tx.Rollback()
}

func (s *sqlTx) Stmt(stmt *sql.Stmt) Stmt {
	stmt1 := s.Tx.Stmt(stmt)
	return &sqlStmt{
		Stmt: stmt1,
	}
}

func (s *sqlTx) StmtContext(ctx context.Context, stmt *sql.Stmt) Stmt {
	stmt1 := s.Tx.StmtContext(ctx, stmt)
	return &sqlStmt{
		Stmt: stmt1,
	}
}
