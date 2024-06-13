.DEFAULT_GOAL := help
.PHONY: help

VERSION=$(shell cat ./VERSION)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test:
	go test ./... -cover

tag: ## Create and push a new tag
	git tag -a ${VERSION} -m "Release ${VERSION}"
	git push origin ${VERSION}
