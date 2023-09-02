/*
 * Copyright (c) 2023. App Nerds LLC. All rights reserved
 */
package workerpool2

import (
	"sync"
)

type WorkerFunc func() error

/*
PoolOrchestrator describes an interface for managing a pool of workers
who perform jobs
*/
type PoolOrchestrator interface {
	PutWorkerInTheQueue(worker Worker)
	Shutdown()
	Start()
	QueueJob(job WorkerFunc)
	Wait()
	WriteError(err error)
}

/*
PoolConfig provides the ability to configure the worker
pool. MaxWorkers specifies the maximum number of workers
available. ErrorChan allows you to provide a channel to
watch for errors. This is optional
*/
type PoolConfig struct {
	ErrorChan  chan error
	MaxWorkers int
}

/*
A Pool provides methods for managing a pool of workers who
perform jobs. A pool can be configured to have a maximum number
of available workers, and will wait up to a configurable amount
of time for a worker to become available before returning an error
*/
type Pool struct {
	activeJobs      *sync.WaitGroup
	errorChan       chan error
	jobQueue        chan WorkerFunc
	maxWorkers      int
	shutdownChannel chan struct{}
	workerQueue     chan Worker
}

/*
NewPool creates a new Pool
*/
func NewPool(config PoolConfig) *Pool {
	return &Pool{
		activeJobs:      &sync.WaitGroup{},
		errorChan:       config.ErrorChan,
		jobQueue:        make(chan WorkerFunc),
		maxWorkers:      config.MaxWorkers,
		shutdownChannel: make(chan struct{}),
		workerQueue:     make(chan Worker, config.MaxWorkers),
	}
}

/*
assignJob attempts to assign a job to a worker, if available. If a
worker is not available an error is returned
*/
func (p *Pool) assignJob(job WorkerFunc) {
	select {
	case worker := <-p.workerQueue:
		if worker != nil {
			worker.DoJob(job)
		}
	}
}

/*
PutWorkerInTheQueue puts a worker in the worker queue
*/
func (p *Pool) PutWorkerInTheQueue(worker Worker) {
	p.workerQueue <- worker
}

/*
Shutdown closes the job queue and waits for current workers to finish
*/
func (p *Pool) Shutdown() {
	p.shutdownChannel <- struct{}{}
}

/*
Start hires workers and waits for jobs
*/
func (p *Pool) Start() {
	for index := 0; index < p.maxWorkers; index++ {
		p.workerQueue <- &PoolWorker{
			pool:      p,
			waitGroup: p.activeJobs,
		}
	}

	go func() {
		for {
			select {
			case <-p.shutdownChannel:
				break

			case job := <-p.jobQueue:
				p.assignJob(job)
			}
		}
	}()
}

/*
QueueJob adds a job to the work queue
*/
func (p *Pool) QueueJob(job WorkerFunc) {
	p.activeJobs.Add(1)
	p.jobQueue <- job
}

/*
Wait waits for active jobs to finish
*/
func (p *Pool) Wait() {
	p.activeJobs.Wait()
}

/*
WriteError writes an error to the error channel, if available.
*/
func (p *Pool) WriteError(err error) {
	if p.errorChan != nil {
		p.errorChan <- err
	}
}
