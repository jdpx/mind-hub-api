GO_BUILD = GOOS=linux go build
FUNCTIONS_LAMBDAS = $(wildcard cmd/lambdas/*/main.go)
FUNCTIONS_DIRS = $(shell ls lambdas)

TERRAFORM_ACTION=plan
APP_ENV=dev
OPTS=

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


.PHONY: run-api
run-api:
	go run ./cmd/api/.

.PHONY: regenerate-types
regenerate-types:
	go run github.com/99designs/gqlgen generate

cmd/lambdas/%/main.go:
	cd $(subst main.go,,$@) \
	&& $(GO_BUILD) -o lambda ./.\
	&& zip ../../../dist/$*.zip lambda \
	&& rm lambda

.PHONY: build-all-lambdas
build-all-lambdas:
	mkdir -p dist
	make $(FUNCTIONS_LAMBDAS)

.PHONY: validate-terraform
validate-terraform:
	cd terraform/providers/aws/$(APP_ENV) && \
	terraform init && \
	terraform validate && \
	terraform $(TERRAFORM_ACTION) $(OPTS)

.PHONY: regenerate-all-mocks
regenerate-all-mocks:
	@go generate ./...