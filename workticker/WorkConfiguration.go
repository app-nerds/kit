package workticker

/*
WorkConfiguration is used to configure how work is retrieved and handled
*/
type WorkConfiguration[T any] struct {
	Handler   WorkHandler[T]
	Retriever WorkRetriever[T]
}
