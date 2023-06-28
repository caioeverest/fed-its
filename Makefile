envfile:=.env.local
ifeq ($(shell test ! -f .env.local && echo -n yes),yes)
    envfile=env.example
endif

include $(envfile)
export $(shell sed 's/=.*//' $(envfile))
export BRANCH=$(shell git branch --show-current | cut -d '/' -f2)
version:=${BRANCH}

lint:
	@golangci-lint run ./src/... --allow-parallel-runners

update-swagger:
	@swag i fmt ./
	@swag i -d ./

update-mocks:
	@mockery

test:
	@go test ./... -v -failfast

build:
	@go build -ldflags "-X internal/config/config.Version=${version}" -o ./bin/fedits ./src/main.go

run:
	@go run ./main.go

watch:
	# increase the file watch limit, might required on MacOS
	ulimit -n 1638400
	air

install-tools:
	@go install github.com/cosmtrek/air@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
