package workticker

import "fmt"

type HandleWorkError[T any] struct {
	ErrorMessage   string
	WorkerID       int
	Data           T
	WorkTickerName string
}

func NewHandleWorkError[T any](errorMessage string, workerID int, data T) HandleWorkError[T] {
	return HandleWorkError[T]{
		ErrorMessage: errorMessage,
		WorkerID:     workerID,
		Data:         data,
	}
}

func (hwe HandleWorkError[T]) Error() string {
	return fmt.Sprintf("[%s] error handling work in worker ID %d: %s", hwe.WorkTickerName, hwe.WorkerID, hwe.ErrorMessage)
}
