package workticker

import "fmt"

type HandleWorkError[T any] struct {
	ErrorMessage string
	WorkerID     int
	Data         T
}

func NewHandleWorkError[T any](errorMessage string, workerID int, data T) HandleWorkError[T] {
	return HandleWorkError[T]{
		ErrorMessage: errorMessage,
		WorkerID:     workerID,
		Data:         data,
	}
}

func (hwe HandleWorkError[T]) Error() string {
	return fmt.Sprintf("error handling work in worker ID %d: %s", hwe.WorkerID, hwe.ErrorMessage)
}
