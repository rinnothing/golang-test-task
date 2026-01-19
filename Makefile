.PHONY: install-deps
install-deps:
	go get -tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: generate-api
generate-api: install-deps
	go tool oapi-codegen -generate=types -package=gen -o api/gen/types.gen.go api/openapi.yml
	go tool oapi-codegen -generate=server -package=gen -o api/gen/server.gen.go api/openapi.yml
	go tool oapi-codegen -generate=client -package=gen -o api/gen/client.gen.go api/openapi.yml
	go mod tidy

.PHONY: build
build: generate-api
	go build -o server cmd/main.go

.PHONY: start
start:
	docker-compose build
	docker-compose up

.PHONY: integration-test
integration-test:
	cd integration && go test ./...

.PHONY: drop-volume
drop-volume:
	docker-compose down
	docker volume rm golang-test-task_postgres_data
