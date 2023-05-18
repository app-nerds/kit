package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Retry(tries, backoff int, fn func(try int) error) error {
	var (
		err     error
		attempt int = 1
	)

	if tries < 1 || backoff < 1 {
		return fmt.Errorf("tries and backoff must be greater than 0")
	}

tryagain:
	if err = fn(attempt); err != nil {
		if attempt >= tries {
			return err
		}

		attempt++
		currentWait := calculateWaitTime(attempt, backoff)
		time.Sleep(time.Second * time.Duration(currentWait))

		goto tryagain
	}

	return err
}

func calculateWaitTime(tries, backoff int) int64 {
	jitter := rand.Intn(backoff)
	return int64(math.Pow(float64(backoff+jitter), float64(tries-1)))
}
