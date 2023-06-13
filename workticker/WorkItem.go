package workticker

/*
WorkItem is a single unit of work.
*/
type WorkItem[T any] struct {
	Data    T
	Handler WorkHandler[T]
}
