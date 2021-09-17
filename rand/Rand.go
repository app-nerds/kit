/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package rand

import (
	"math/rand"
	"time"
)

/*
String generates a random alpha-numeric string
*/
func String(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range b {
		b[index] = characters[seed.Intn(len(characters))]
	}

	return string(b)
}
