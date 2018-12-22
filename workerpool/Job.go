// Copyright 2018-2019 AppNinjas LLC. All rights reserved
// Use of this source code is governed by the MIT license.
package workerpool

/*
A Job is an interface structs must implement which actually executes
the work to be done by a worker in the pool. The workerID is the
identifier of the worker performing the job
*/
type Job interface {
	Work(workerID int)
}
