# Worker Pool v2

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

	"github.com/app-nerds/kit/v6/workerpool2"
)

func newWorker(index int) func() error {
    return func() error {
        fmt.Printf("Worker %d sleeping...\n", index)
        time.Sleep(2 * time.Second)
    }
}

func main() {
	var pool workerpool2.PoolOrchestrator

    errChan := make(chan error)
    stopErrChan := make(chan struct{})

    go func() {
        for {
            select {
            case err := <-errChan:
                // Do something useful with any errors
                fmt.Printf("we received error '%v'\n", err)

            case <-stopErrChan:
                return
            }
        }
    }()

	pool = workerpool2.NewPool(workerpool2.PoolConfig{
        ErrorChan:  errChan,
		MaxWorkers: 10,
	})

	pool.Start()

	for index := 0; index < 30; index++ {
        job := newWorker(index)
        pool.QueueJob(job)
	}

	pool.Wait()
	pool.Shutdown()
    stopErrChan <- struct{}{}
}
```

