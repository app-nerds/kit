package mongocertstore

import (
	"context"
	"time"

	"github.com/app-nerds/kit/v4/database"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/acme/autocert"
)

/*
CertCache implements the autocert Cache interface and provides the ability
to store SSL certs in the MongoDB. TODO: Add ability to mock Collection
*/
type CertCache struct {
	Collection database.Collection
	DB         database.Database
}

/*
NewCertCache creates a new SSL certificate cache in a MongoDB collection
*/
func NewCertCache(db database.Database, collectionName string) *CertCache {
	return &CertCache{
		Collection: db.C(collectionName),
		DB:         db,
	}
}

/*
Delete removes a certificate from the database. The cert to remove
is identified by a key
*/
func (cc *CertCache) Delete(ctx context.Context, key string) error {
	selector := bson.M{
		"key": key,
	}

	cc.Collection.Remove(selector)
	return nil
}

/*
Get retrieves a certificate by key
*/
func (cc *CertCache) Get(ctx context.Context, key string) ([]byte, error) {
	var err error
	var result *CertCacheItem

	selector := bson.M{
		"key": key,
	}

	if err = cc.Collection.Find(selector).One(&result); err != nil {
		if err == mgo.ErrNotFound {
			return nil, autocert.ErrCacheMiss
		}

		return nil, err
	}

	return result.Certificate, nil
}

/*
Put inserts a certificate into the database
*/
func (cc *CertCache) Put(ctx context.Context, key string, data []byte) error {
	certificate := &CertCacheItem{
		Certificate:        data,
		DateTimeCreatedUTC: time.Now().UTC(),
		Key:                key,
	}

	return cc.Collection.Insert(certificate)
}
