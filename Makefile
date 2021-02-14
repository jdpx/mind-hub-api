BUILD_NUMBER = 0.0.0
GO_BUILD = GOOS=linux go build -ldflags="-X main.buildVersion=$(BUILD_NUMBER)"
FUNCTIONS_LAMBDAS = $(wildcard cmd/lambdas/*/main.go)
FUNCTIONS_DIRS = $(shell ls lambdas)

TERRAFORM_ACTION=plan
APP_ENV=dev
OPTS=

.PHONY: install
install:
	go mod download

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build/lambdas
build/lambdas:
	mkdir -p dist
	make $(FUNCTIONS_LAMBDAS)

cmd/lambdas/%/main.go:
	cd $(subst main.go,,$@) \
	&& $(GO_BUILD) -o lambda ./.\
	&& zip ../../../dist/$*.zip lambda \
	&& rm lambda

.PHONY: validate-terraform
validate-terraform:
	cd terraform/providers/aws/$(APP_ENV) && \
	terraform init && \
	terraform validate && \
	terraform $(TERRAFORM_ACTION) $(OPTS)

.PHONY: run/tfsec
run/tfsec:
	tfsec . -e AWS002,AWS017

.PHONY: run/api
run/api:
	go run ./cmd/api/.

.PHONY: run/local-dynamo
run/local-dynamo:
	docker-compose up

.PHONY: run/tests
run/tests:
	go test ./... -cover

.PHONY: run/lint
run/lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
	golangci-lint run

.PHONY: generate/graphql
generate/graphql:
	go run github.com/99designs/gqlgen generate

.PHONY: generate/mocks
generate/mocks:
	@go generate ./...

