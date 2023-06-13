package workerpoolv2

import (
	"go.uber.org/ratelimit"
)

type WorkItem struct {
	Data    interface{}
	Handler WorkHandler
}

type WorkItemCollection []WorkItem

func (wic WorkItemCollection) Handle(workerID int, limiter ratelimit.Limiter, handler WorkHandler) error {
	work := make([]interface{}, len(wic))

	for i, item := range wic {
		work[i] = item.Data
	}

	err := handler.ProcessWork(workerID, work, limiter)

	if err != nil {
		return NewHandleWorkError(err.Error(), workerID, work)
	}

	return nil
}
