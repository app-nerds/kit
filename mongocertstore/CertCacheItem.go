/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package mongocertstore

import (
	"time"
)

type CertCacheItem struct {
	Certificate        []byte    `json:"certificate" bson:"certificate"`
	DateTimeCreatedUTC time.Time `json:"dateTimeCreatedUTC" bson:"dateTimeCreatedUTC"`
	Key                string    `json:"key" bson:"key"`
}
