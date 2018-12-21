package workerpool

import (
	"time"
)

/*
PoolConfig provides the ability to configure the worker
pool. MaxWorkers specifies the maximum number of workers
available. This essentially sets the channel size.
MaxWorkerWaitTime is a duration that tells the pool how
long it will wait before timing out when a client requests
a worker.
*/
type PoolConfig struct {
	MaxJobQueue       int
	MaxWorkers        int
	MaxWorkerWaitTime time.Duration
}
