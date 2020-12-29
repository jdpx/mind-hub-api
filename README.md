# Mind Hub API

## Setup

- Clone the project from Github
  - `git clone git@github.com:jdpx/mind-hub-api.git`
- install Golang
  - `brew install go`
- Install dependancies
  - `make install`
- Install [Docker](https://docs.docker.com/get-docker/)
- Install Golang-Lint
  - `brew install golangci-lint`

## Development

- Start local DynamoDB
  - `make run/local-dynamo`
- Run local API
  - `make run/api`

## Commands

- Run tests
  - `make run/tests`
- Run linter
  - `make lint`
- Regenerate GraphQL Schema
  - `make generate/graphql`
