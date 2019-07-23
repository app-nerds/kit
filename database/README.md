# Database

A set of interface wrappers for Go's MongoDB driver

## Usage

Usage is just like using the [Mongo driver](https://github.com/globalsign/mgo). The biggest
difference is that instead of each object being a struct you will work with interfaces
by the same name. This allow you to mock and test your code. Here is a small example
usage.

```go
package main

import (
	"github.com/app-nerds/kit/database"
	"github.com/globalsign/mgo/bson"
)

func main() {
	var err error
	var session database.Session
	var testData []string

	if session, err = database.Dial("localhost:27017"); err != nil {
		panic(err)
	}

	db := session.DB("mydatabase")
	defer session.Close()

	collection := db.C("collection")

	if err = collection.Find(bson.M{}).All(&testData); err != nil {
		panic(err)
	}
}
```

## Testing

This package provides a full set of mock structs that allow you to mock
your database interactions for unit testing. Below is an example.

```go
package example_test

import (
	"reflect"
	"testing"

	"pkg/example"
	"github.com/app-nerds/kit/database"
)

func TestSomeFunc(t *testing.T) {
	allCalled := false
	findCalled := false

	expected := []string{
		"value 1",
		"value 2",
	}

	mock := &database.CollectionMock{
		FindFunc: func(query interface{}) database.Query {
			findCalled = true

			return &database.QueryMock{
				AllFunc: func(result interface{}) error {
					allCalled = true
					database.WriteToResultInterface([]string{"value 1", "value 2"})
					return nil
				},
			}
		},
	}

	component := example.Example{
		Collection: mock,
	}

	actual, err := component.SomeMethod()

	if err != nil {
		t.Errorf("Expected err to be nil")
	}

	if !allCalled {
		t.Errorf("Expected All() to be called")
	}

	if !findCalled {
		t.Errorf("Expected Find() to be called")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v to be %+v", actual, expected)
	}
}
```

## License

MIT License

Copyright (c) 2019 App Nerds LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
