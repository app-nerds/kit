package sqldatabase

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

var (
	KindString         reflect.Kind = reflect.String
	KindInt            reflect.Kind = reflect.Int64
	KindBool           reflect.Kind = reflect.Bool
	KindFloat32        reflect.Kind = reflect.Float32
	KindFloat64        reflect.Kind = reflect.Float64
	KindTime           reflect.Kind = reflect.TypeOf(time.Time{}).Kind()
	KindSqlNullString  reflect.Kind = reflect.TypeOf(sql.NullString{}).Kind()
	KindSqlNullInt32   reflect.Kind = reflect.TypeOf(sql.NullInt32{}).Kind()
	KindSqlNullInt64   reflect.Kind = reflect.TypeOf(sql.NullInt64{}).Kind()
	KindSqlNullTime    reflect.Kind = reflect.TypeOf(sql.NullTime{}).Kind()
	KindSqlNullFloat64 reflect.Kind = reflect.TypeOf(sql.NullFloat64{}).Kind()
)

/*
ScanMappingItem decribes a value and what kind of value it is.

Example 1: A string value

  mapping := gosqldatabase.ScanMappingItem{gosqldatabase.KindString, "value"}

Example 2: A sql.NullInt64 value

  mapping := gosqldatabase.ScanMappingItem{gosqldatabase.KindSqlNullInt64, sql.NullInt64{25, true}}
*/
type ScanMappingItem struct {
	Kind  reflect.Kind
	Value interface{}
}

/*
ScanMapping is a slice of a slice of ScanMappingItem structs. Think of this
as a set of rows containing a set of column definitions.

Example:

  data := gosqldatabase.ScanMapping{
    {
		 {gosqldatabase.KindString, "value1"},
		 {gosqldatabase.KindInt, 2},
		 {gosqldatabase.KindSqlNullString, sql.NullString{"value2", true}},
		 {gosqldatabase.KindSqlNullInt64, sql.NullInt64{nil, false}},
    },
  }
*/
type ScanMapping [][]ScanMappingItem

/*
Scan scans values into dest... from the provided mappings and row index.
This is most useful in a mock method.

Example:

  data := gosqldatabase.ScanMapping{
    {
		 {gosqldatabase.KindString, "value1"},
		 {gosqldatabase.KindInt, 2},
		 {gosqldatabase.KindSqlNullString, sql.NullString{"value2", true}},
		 {gosqldatabase.KindSqlNullInt64, sql.NullInt64{nil, false}},
    },
  }

  rowIndex := 0

  mock := &gosqldatabase.MockRow{
    ScanFunc: func(dest ...interface{}) error {
      gosqldatabase.Scan(data, rowIndex, dest...)
      return nil
    },
  }
*/
func Scan(mappings ScanMapping, rowIndex int, dest ...interface{}) {
	for colIndex, d := range dest {
		AssignScanValue(mappings, rowIndex, colIndex, d)
	}
}

/*
AssignScanValue reads the mapping at a row and column index, determines
the type of value, and assigns it to the provided destination variable.
*/
func AssignScanValue(mappings ScanMapping, rowIndex, colIndex int, dest interface{}) {
	var ok bool

	wrongType := func(rowIndex, colIndex int, expectedType string) {
		fmt.Printf("value at row %d, col %d is not %s\n", rowIndex, colIndex, expectedType)
	}

	switch mappings[rowIndex][colIndex].Kind {
	case reflect.String:
		var value string
		p := dest.(*string)

		if value, ok = mappings[rowIndex][colIndex].Value.(string); !ok {
			wrongType(rowIndex, colIndex, "string")
			return
		}

		*p = value

	case reflect.Int16, reflect.Int64:
		var value int
		p := dest.(*int)

		if value, ok = mappings[rowIndex][colIndex].Value.(int); !ok {
			wrongType(rowIndex, colIndex, "int")
			return
		}

		*p = value

	case reflect.Float32:
		var value float32
		p := dest.(*float32)

		if value, ok = mappings[rowIndex][colIndex].Value.(float32); !ok {
			wrongType(rowIndex, colIndex, "float32")
			return
		}

		*p = value

	case reflect.Float64:
		var value float64
		p := dest.(*float64)

		if value, ok = mappings[rowIndex][colIndex].Value.(float64); !ok {
			wrongType(rowIndex, colIndex, "float64")
			return
		}

		*p = value

	case reflect.Bool:
		var value bool
		p := dest.(*bool)

		if value, ok = mappings[rowIndex][colIndex].Value.(bool); !ok {
			wrongType(rowIndex, colIndex, "bool")
			return
		}

		*p = value

	default:
		switch reflect.TypeOf(mappings[rowIndex][colIndex].Value).String() {
		case "time.Time":
			var value time.Time
			p := dest.(*time.Time)

			if value, ok = mappings[rowIndex][colIndex].Value.(time.Time); !ok {
				wrongType(rowIndex, colIndex, "time.Time")
				return
			}

			*p = value

		case "sql.NullString":
			var value sql.NullString
			p := dest.(*sql.NullString)

			if value, ok = mappings[rowIndex][colIndex].Value.(sql.NullString); !ok {
				wrongType(rowIndex, colIndex, "sql.NullString")
				return
			}

			*p = value

		case "sql.NullInt32":
			var value sql.NullInt32
			p := dest.(*sql.NullInt32)

			if value, ok = mappings[rowIndex][colIndex].Value.(sql.NullInt32); !ok {
				wrongType(rowIndex, colIndex, "sql.NullInt32")
				return
			}

			*p = value

		case "sql.NullInt64":
			var value sql.NullInt64
			p := dest.(*sql.NullInt64)

			if value, ok = mappings[rowIndex][colIndex].Value.(sql.NullInt64); !ok {
				wrongType(rowIndex, colIndex, "sql.NullInt64")
				return
			}

			*p = value

		case "sql.NullTime":
			var value sql.NullTime
			p := dest.(*sql.NullTime)

			if value, ok = mappings[rowIndex][colIndex].Value.(sql.NullTime); !ok {
				wrongType(rowIndex, colIndex, "sql.NullTime")
				return
			}

			*p = value

		case "sql.NullFloat64":
			var value sql.NullFloat64
			p := dest.(*sql.NullFloat64)

			if value, ok = mappings[rowIndex][colIndex].Value.(sql.NullFloat64); !ok {
				wrongType(rowIndex, colIndex, "sql.NullFloat64")
				return
			}

			*p = value
		}
	}
}
