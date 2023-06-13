package workticker

import (
	"go.uber.org/ratelimit"
)

/*
WorkHandler is a function that handles a single unit of work. It should do whatever processing is necessary.
If an error is returned, that error is placed onto the error channel.
*/
type WorkHandler[T any] func(workerID int, data T, limiter ratelimit.Limiter) error
