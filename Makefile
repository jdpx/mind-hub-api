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

cmd/lambdas/%/main.go:
	echo goo
	cd $(subst main.go,,$@) \
	&& $(GO_BUILD) -o lambda ./.\
	&& zip ../../../dist/$*.zip lambda \
	&& rm lambda

.PHONY: build-all-lambdas
build-all-lambdas: 
	make $(FUNCTIONS_LAMBDAS)