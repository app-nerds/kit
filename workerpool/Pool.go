/*
Copyright 2018-2019 AppNinjas LLC. All rights reserved
Use of this source code is governed by the MIT license.

Package workerpool implements a pool of workers (goroutines) who
can perform jobs. Below is an example of a pool that has
10 workers, and if none are available will wait 2 seoncds
for a worker to become available before returning a timeout error.

	package main

	import (
		"fmt"
		"time"

		"code.appninjas.biz/appninjas/kit/workerpool"
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
*/
package workerpool

import (
	"sync"
	"time"
)

/*
IPool describes an interface for managing a pool of workers
who perform jobs
*/
type IPool interface {
	AssignJob(job Job) JobError
	PutWorkerInTheQueue(worker IWorker)
	Shutdown()
	Start()
	QueueJob(job Job)
	Wait()
}

/*
A Pool provides methods for managing a pool of workers who
perform jobs. A pool can be configured to have a maximum number
of available workers, and will wait up to a configurable amount
of time for a worker to become available before returning an error
*/
type Pool struct {
	activeJobs      *sync.WaitGroup
	config          PoolConfig
	jobQueue        chan Job
	shutdownChannel chan bool
	workerQueue     chan IWorker

	ErrorQueue chan JobError
}

/*
ErrNoAvaialableWorkers is used to describe a situation where there
are no workers available to work a job
*/
type ErrNoAvaialableWorkers struct {
	Job Job
}

func (e ErrNoAvaialableWorkers) Error() string {
	return "No workers available"
}

/*
GetJob returns the job associated with this error
*/
func (e ErrNoAvaialableWorkers) GetJob() Job {
	return e.Job
}

/*
NewPool creates a new Pool
*/
func NewPool(config PoolConfig) *Pool {
	return &Pool{
		activeJobs:      &sync.WaitGroup{},
		config:          config,
		jobQueue:        make(chan Job, config.MaxJobQueue),
		shutdownChannel: make(chan bool),
		workerQueue:     make(chan IWorker, config.MaxWorkers),

		ErrorQueue: make(chan JobError),
	}
}

/*
AssignJob attempts to assign a job to a worker, if available. If a
worker is not available an error is returned
*/
func (p *Pool) AssignJob(job Job) JobError {
	select {
	case worker := <-p.workerQueue:
		if worker != nil {
			worker.DoJob(job)
		}

		return nil

	case <-time.After(p.config.MaxWorkerWaitTime):
		return ErrNoAvaialableWorkers{Job: job}
	}
}

/*
PutWorkerInTheQueue puts a worker in the worker queue
*/
func (p *Pool) PutWorkerInTheQueue(worker IWorker) {
	p.workerQueue <- worker
}

/*
Shutdown closes the job queue and waits for current workers to finish
*/
func (p *Pool) Shutdown() {
	p.shutdownChannel <- true
}

/*
Start hires workers and waits for jobs
*/
func (p *Pool) Start() {
	for index := 1; index <= p.config.MaxWorkers; index++ {
		p.workerQueue <- &Worker{
			Pool:      p,
			WaitGroup: p.activeJobs,
			WorkerID:  index,
		}
	}

	go func() {
		for {
			select {
			case <-p.shutdownChannel:
				break

			case job := <-p.jobQueue:
				err := p.AssignJob(job)

				if err != nil {
					p.ErrorQueue <- err
				}
			}
		}
	}()
}

/*
QueueJob adds a job to the work queue
*/
func (p *Pool) QueueJob(job Job) {
	p.activeJobs.Add(1)
	p.jobQueue <- job
}

/*
Wait waits for active jobs to finish
*/
func (p *Pool) Wait() {
	p.activeJobs.Wait()
}
