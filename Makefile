DOCKER_VERSION?=1

test: ## Runs tests
	go test ./...
build: ## Builds the binary
	go build .
linux: ## Builds linux binary
	GOOS=linux GOARCH=amd64 go build -o="dist/linux/dkr" && \
	tar -czvf dist/dkr.linux.amd64.tar.gz dist/linux
osx: ## Builds OSx binary
	GOOS=darwin GOARCH=amd64 go build -o="dist/osx/dkr" && \
	tar -czvf dist/dkr.darwin.amd64.tar.gz dist/osx
windows: ## Builds windows binary
	GOOS=windows GOARCH=amd64 go build -o="dist/windows/sm.exe" && \
	zip dist/dkr.windows.amd64.zip dist/windows/sm.exe
binaries: ## Builds all binaries
	$(MAKE) linux && $(MAKE) osx && $(MAKE) windows
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
