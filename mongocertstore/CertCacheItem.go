package mongocertstore

import (
	"time"
)

type CertCacheItem struct {
	Certificate        []byte    `json:"certificate" bson:"certificate"`
	DateTimeCreatedUTC time.Time `json:"dateTimeCreatedUTC" bson:"dateTimeCreatedUTC"`
	Key                string    `json:"key" bson:"key"`
}
