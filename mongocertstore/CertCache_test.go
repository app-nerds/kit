/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package mongocertstore_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/app-nerds/kit/v6/database"
	"github.com/app-nerds/kit/v6/mongocertstore"
	"github.com/globalsign/mgo"
	"golang.org/x/crypto/acme/autocert"
)

func TestNewCertCache(t *testing.T) {
	db := &database.DatabaseMock{
		CFunc: func(name string) database.Collection {
			return nil
		},
	}

	actual := mongocertstore.NewCertCache(db, "certcache")

	isCertCache := func(i interface{}) bool {
		switch i.(type) {
		case *mongocertstore.CertCache:
			return true
		default:
			return false
		}
	}

	if !isCertCache(actual) {
		t.Error("Expected a CertCache object")
	}
}

func TestDelete(t *testing.T) {
	removeFuncCalled := false

	cache := &mongocertstore.CertCache{
		Collection: &database.CollectionMock{
			RemoveFunc: func(selector interface{}) error {
				removeFuncCalled = true
				return nil
			},
		},
	}

	ctx := context.Background()

	actual := cache.Delete(ctx, "key")

	if actual != nil {
		t.Errorf("Expected nil error")
	}

	if !removeFuncCalled {
		t.Errorf("Expected the remove function to be called")
	}
}

func TestGet(t *testing.T) {
	findFuncCalled := false
	oneFuncCalled := false
	expected := []byte("bytes")

	queryMock := &database.QueryMock{
		OneFunc: func(result interface{}) error {
			oneFuncCalled = true

			v := &mongocertstore.CertCacheItem{
				Certificate: expected,
			}

			return database.WriteToResultInterface(v, result)
		},
	}

	collectionMock := &database.CollectionMock{
		FindFunc: func(query interface{}) database.Query {
			findFuncCalled = true
			return queryMock
		},
	}

	certCache := &mongocertstore.CertCache{
		Collection: collectionMock,
	}

	ctx := context.Background()

	actual, err := certCache.Get(ctx, "key")

	if err != nil {
		t.Errorf("Expected err to be nil")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected certificate bytes to match")
	}

	if !findFuncCalled {
		t.Errorf("Expected the Find function to be called")
	}

	if !oneFuncCalled {
		t.Errorf("Expected the One function to be called")
	}
}

func TestGetNotFoundError(t *testing.T) {
	findFuncCalled := false
	oneFuncCalled := false

	queryMock := &database.QueryMock{
		OneFunc: func(result interface{}) error {
			oneFuncCalled = true
			return mgo.ErrNotFound
		},
	}

	collectionMock := &database.CollectionMock{
		FindFunc: func(query interface{}) database.Query {
			findFuncCalled = true
			return queryMock
		},
	}

	certCache := &mongocertstore.CertCache{
		Collection: collectionMock,
	}

	ctx := context.Background()
	actual, err := certCache.Get(ctx, "key")

	if err != autocert.ErrCacheMiss {
		t.Errorf("Expected error to be autocert.ErrCacheMiss")
	}

	if actual != nil {
		t.Errorf("Expected certificate to be nil")
	}

	if !findFuncCalled {
		t.Errorf("Expected the Find function to be called")
	}

	if !oneFuncCalled {
		t.Errorf("Expected the One function to be called")
	}
}

func TestGetDatabaseError(t *testing.T) {
	findFuncCalled := false
	oneFuncCalled := false
	expected := fmt.Errorf("Error in database")

	queryMock := &database.QueryMock{
		OneFunc: func(result interface{}) error {
			oneFuncCalled = true
			return expected
		},
	}

	collectionMock := &database.CollectionMock{
		FindFunc: func(query interface{}) database.Query {
			findFuncCalled = true
			return queryMock
		},
	}

	certCache := &mongocertstore.CertCache{
		Collection: collectionMock,
	}

	ctx := context.Background()
	actual, err := certCache.Get(ctx, "key")

	if err != expected {
		t.Errorf("Expected error to be Error in database")
	}

	if actual != nil {
		t.Errorf("Expected certificate to be nil")
	}

	if !findFuncCalled {
		t.Errorf("Expected the Find function to be called")
	}

	if !oneFuncCalled {
		t.Errorf("Expected the One function to be called")
	}
}

func TestPut(t *testing.T) {
	insertFuncCalled := false
	expectedKey := "Key"
	expectedData := []byte("data")
	actualKey := ""
	actualData := []byte("")

	collectionMock := &database.CollectionMock{
		InsertFunc: func(docs ...interface{}) error {
			insertFuncCalled = true
			c := docs[0].(*mongocertstore.CertCacheItem)

			actualKey = c.Key
			actualData = c.Certificate

			return nil
		},
	}

	certCache := &mongocertstore.CertCache{
		Collection: collectionMock,
	}

	ctx := context.Background()
	err := certCache.Put(ctx, expectedKey, expectedData)

	if !insertFuncCalled {
		t.Errorf("Expected insert function to be called")
	}

	if err != nil {
		t.Errorf("Expect result to be nil")
	}

	if actualKey != expectedKey {
		t.Errorf("Expected keys to match")
	}

	if !reflect.DeepEqual(expectedData, actualData) {
		t.Errorf("Expected certificate data to match")
	}
}
