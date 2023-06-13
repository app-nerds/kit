package workerpoolv2

type WorkRetriever interface {
	RetrieveWork(handler WorkHandler) ([]WorkItem, error)
}
