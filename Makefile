.PHONY:

build:
	docker compose up --build -d

compose-up:
	docker compose up --build -d postgres rabbitmq && docker-compose logs -f

up:
	docker compose up -d && docker compose logs -f
compose-down:
	docker compose down --remove-orphans

.PHONY: tester
tester:
	docker compose up -f docker-compose.tester.yml -d

test:
	go test -v -cover ./internal/... ./pkg/... ./cmd/...
.PHONY: test

linter-golangci:
	golangci-lint run

.PHONY: proto
proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
		protoc --proto_path=proto --go_out=proto/grpc --go_opt=paths=source_relative   --go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative   proto/*.proto; \

clean:
	go clean -i google.golang.org/grpc/...
