package workticker

/*
WorkRetriever is a function that returns a WorkItem and an error. It is used to retrieve work items on a tick.
*/
type WorkRetriever[T any] func(handler WorkHandler[T]) (WorkItem[T], error)
