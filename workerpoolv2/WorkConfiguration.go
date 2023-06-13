package workerpoolv2

type WorkConfiguration struct {
	Handler    WorkHandler
	MaxRetries int
	Retriever  WorkRetriever
}
