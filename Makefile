.PHONY: init mocks docs unit clean devrun cover info.out build

DEV_RUN_PORT := 13434

init: docs mocks
	@echo '==================== make init ====================='
	go mod tidy

mocks:
	@echo '==================== make mocks ====================='
	mockery --all --keeptree

docs:
	@echo '==================== make docs ====================='
	swag init --generalInfo cmd/main.go -o docs

unit:
	@echo '==================== make unit ====================='
	go test ./... -v

clean:
	@echo '==================== make clean ====================='
	go clean -cache
	rm -rf mocks logs docs info.out coverage*.out service

devrun:
	@echo '==================== make devrun ====================='
	 GIN_MODE=debug go run ./cmd/main.go -port=$(DEV_RUN_PORT) -config=./config/dev/application.yaml

cover:
	@echo '==================== make cover ====================='
	go test -coverprofile=coverage.out ./...;go tool cover -html=coverage.out

info.out:
	@echo '==================== make info.out ====================='
	./get_build_info.sh

build: info.out
	@echo '==================== make build ====================='
	go build -tags=jsoniter -ldflags='-X "main.BuildInfo=$(shell cat info.out)"' -o service cmd/main.go
