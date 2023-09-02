/*
 * Copyright (c) 2023. App Nerds LLC. All rights reserved
 */
package workerpool2

import "sync"

/*
Worker interface describes a struct that performs a job
*/
type Worker interface {
	DoJob(job WorkerFunc)
	RejoinWorkerPool()
}

/*
A PoolWorker is someone that performs a job. There are a finite number
of workers in the pool
*/
type PoolWorker struct {
	pool      PoolOrchestrator
	waitGroup *sync.WaitGroup
}

/*
DoJob executes the provided job. When the work is complete
this worker will put itself back in the queue as available.
This method execute a goroutine
*/
func (w *PoolWorker) DoJob(job WorkerFunc) {
	go func(j WorkerFunc) {
		err := j()

		if err != nil {
			w.pool.WriteError(err)
		}

		w.RejoinWorkerPool()
		w.waitGroup.Done()
	}(job)
}

/*
RejoinWorkerPool puts this worker back in the worker queue
of the pool. A worker will rejoin the queue when she has
finished the job
*/
func (w *PoolWorker) RejoinWorkerPool() {
	w.pool.PutWorkerInTheQueue(w)
}
