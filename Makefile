ifneq (,$(wildcard ./.env))
	include .env
	export
endif

fileName=scheduler
ldflags += -X 'github.com/wilmacedo/nexco-scheduler/config.INTERVAL="${INTERVAL}"'
ldflags += -s -w

all: build

build:
	@go build -ldflags="${ldflags}" -v -o bin/${fileName} cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test ./ ...