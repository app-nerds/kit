# Worker Pool

A worker pool allows you to setup a pool of finite worker goroutines to perform jobs.
This paradigm is especially useful where you need to limit work in progress. For example,
if you need to run concurrent HTTP requests, you should limit the number of simultaneous
requests because the OS may have limits, and the receieving server may also have limits.

## Example

```golang
package main

import (
	"fmt"
	"time"

	"github.com/app-nerds/kit/workerpool"
)

type Job struct {
	Index int
}

func (j *Job) Work(workerID int) {
	fmt.Printf("Worker %d sleeping on index %d...\n", workerID, j.Index)
	time.Sleep(2 * time.Second)
}

func main() {
	var pool workerpool.IPool

	pool = workerpool.NewPool(workerpool.PoolConfig{
		MaxJobQueue:       100,
		MaxWorkers:        10,
		MaxWorkerWaitTime: 3 * time.Second,
	})

	pool.Start()

	for index := 0; index < 30; index++ {
		job := &Job{Index: index}
		pool.QueueJob(job)
	}

	pool.Wait()
	pool.Shutdown()
}
```
