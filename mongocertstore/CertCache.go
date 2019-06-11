package mongocertstore

import (
	"context"
	"time"

	"github.com/app-nerds/kit/v3/database"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/acme/autocert"
)

/*
CertCache implements the autocert Cache interface and provides the ability
to store SSL certs in the MongoDB. TODO: Add ability to mock Collection
*/
type CertCache struct {
	Collection *mgo.Collection
	DB         database.DocumentDatabase
}

func NewCertCache(db database.DocumentDatabase, collectionName string) *CertCache {
	return &CertCache{
		Collection: db.GetCollection(collectionName),
		DB:         db,
	}
}

func (cc *CertCache) Delete(ctx context.Context, key string) error {
	selector := bson.M{
		"key": key,
	}

	cc.Collection.Remove(selector)
	return nil
}

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

func (cc *CertCache) Put(ctx context.Context, key string, data []byte) error {
	certificate := &CertCacheItem{
		Certificate:        data,
		DateTimeCreatedUTC: time.Now().UTC(),
		Key:                key,
	}

	return cc.Collection.Insert(certificate)
}
