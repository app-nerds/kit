/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package database_test

import (
	"testing"

	"github.com/app-nerds/kit/v6/database"
)

var address string = "localhost:27017"

func setup() {
	s := database.Dial(address)
	db := s.DB("test_kit_database")
	collection := db.C("test_collection")
	collection.DropCollection()
}

// Note that this test requires an active, running MongoDB instance
func TestDial(t *testing.T) {
	var err error
	var session database.Session
	var ok bool

	if session, err = database.Dial(address); err != nil {
		t.Errorf("Error connecting to MongoDB. Notice that this test method requires an active MongoDB installation at %s", address)
	}

	if _, ok = session.(database.Session); !ok {
		t.Errorf("Expected result to be of type Session")
	}
}

func TestDB(t *testing.T) {
	var ok bool

	s, _ := database.Dial(address)
	db := s.DB("test_kit_database")

	if _, ok = db.(database.Database); !ok {
		t.Errorf("Expected result to be of type Database")
	}
}
