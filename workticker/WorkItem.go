package workticker

import "errors"

var (
	ErrNoWorkToRetrieve = errors.New("no work to retrieve")
)

/*
WorkItem is a single unit of work.
*/
type WorkItem[T any] struct {
	Data    T
	Handler WorkHandler[T]
}
