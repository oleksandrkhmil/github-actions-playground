.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -Eh '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run:
	go run ./cmd/main

.PHONY: test
test:
	go test -v -race ./...
