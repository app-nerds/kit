package workerpoolv2

import (
	"fmt"

	"go.uber.org/ratelimit"
)

type WorkHandler interface {
	ProcessWork(workerID int, data []interface{}, limiter ratelimit.Limiter) error
}

type HandleWorkError struct {
	ErrorMessage string
	WorkerID     int
	Data         []interface{}
}

func NewHandleWorkError(errorMessage string, workerID int, data []interface{}) HandleWorkError {
	return HandleWorkError{
		ErrorMessage: errorMessage,
		WorkerID:     workerID,
		Data:         data,
	}
}

func (hwe HandleWorkError) Error() string {
	return fmt.Sprintf("error handling work in worker ID %d: %s", hwe.WorkerID, hwe.ErrorMessage)
}
