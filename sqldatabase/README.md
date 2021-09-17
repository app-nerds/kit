# SQL Database

This package provides interfaces to abstract working with the Go "database/sql" standard library. This allows services you build to unit test SQL interactions.

This README won't go into the details of each interface, as their concrete types are well documented in the Go standard library.

* ColumnType
* DB
* Row
* Rows
* Scanner (this is not in the sdtlib. If abstracts the concent of \*sql.Row and \*sql.Rows for scanning rows)
* Stmt
* Tx

## Examples

### Connect to Database

```go
var (
	err error
	db sqldatabase.DB
)

if db, err = sqldatabase.Open("mysql", "localhost:3306"); err != nil {
	panic(err)
}
```

### Query

```go
func getData() ([]DataStruct, error) {
	var (
		err error
		ctx context.Context
		cancel context.CancelFunc
		rows sqldatabase.Rows
		row DataStruct
	)

	result := make([]DataStruct, 0, 100)
	ctx, cancel = context.WithTimeout(context.Background, time.Second*30)
	defer cancel()

	if rows, err = db.Query("SELECT * FROM stuff"); err != nil {
		return result, err
	}

	for rows.Next() {
		if row, err = convertToStruct(rows); err != nil {
			return result, err
		}

		result = append(result, row)
	}

	return result, nil
}

func convertToStruct(row sqldatabase.Scanner) (DataStruct, error) {
	var (
		err error
		value1 string
		value2 int
		result DataStruct
	)

	if err = row.Scan(&value1, &value2); err != nil {
		return result, err
	}

	result = DataStruct{
		Value1: value1,
		Value2: value2,
	}

	return result, nil
}
```

## DDL

```go
func update() error {
	var (
		err error
		ctx context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if _, err = db.ExecContext(ctx, "UPDATE stuff set a=?", 2); err != nil {
		return err
	}

	return nil
}
```

