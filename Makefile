GO_BUILD = GOOS=linux go build
FUNCTIONS_LAMBDAS = $(wildcard cmd/lambdas/*/main.go)
FUNCTIONS_DIRS = $(shell ls lambdas)

.PHONY: install
install:
	go mod download

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: cb-lint
cb-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
	golangci-lint run

cmd/lambdas/%/main.go:
	cd $(subst main.go,,$@) \
	&& $(GO_BUILD) -o lambda ./.\
	&& zip ../../../dist/$*.zip lambda \
	&& rm lambda

.PHONY: build-all-lambdas
build-all-lambdas:
	mkdir -p dist
	make $(FUNCTIONS_LAMBDAS)