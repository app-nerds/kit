/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */
package workerpool

/*
JobError is an interface to describe an error that has job
information attached
*/
type JobError interface {
	Error() string
	GetJob() Job
}
