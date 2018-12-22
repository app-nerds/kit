// Copyright 2018-2019 AppNinjas LLC. All rights reserved
// Use of this source code is governed by the MIT license.
package workerpool

/*
JobError is an interface to describe an error that has job
information attached
*/
type JobError interface {
	Error() string
	GetJob() Job
}
