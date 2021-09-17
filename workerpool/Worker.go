/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */
package workerpool

import "sync"

/*
IWorker interface describes a struct that performs a job
*/
type IWorker interface {
	DoJob(job Job)
	GetID() int
	RejoinWorkerPool()
}

/*
A Worker is someone that performs a job. There are a finite number
of workers in the pool
*/
type Worker struct {
	Pool      IPool
	WaitGroup *sync.WaitGroup
	WorkerID  int
}

/*
DoJob executes the provided job. When the work is complete
this worker will put itself back in the queue as available.
This method execute a goroutine
*/
func (w *Worker) DoJob(job Job) {
	go func(j Job) {
		j.Work(w.WorkerID)
		w.RejoinWorkerPool()
		w.WaitGroup.Done()
	}(job)
}

/*
GetID returns this worker's ID
*/
func (w *Worker) GetID() int {
	return w.WorkerID
}

/*
RejoinWorkerPool puts this worker back in the worker queue
of the pool. A worker will rejoin the queue when she has
finished the job
*/
func (w *Worker) RejoinWorkerPool() {
	w.Pool.PutWorkerInTheQueue(w)
}
