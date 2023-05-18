all: test testrace

compose-up:
	docker-compose up --build -d postgres rabbitmq && docker-compose logs -f

compose-down:
	docker-compose down --remove-orphans

build:
	go build -o server ./cmd/server/main.go

test: ### run test
	go test -v -cover -race ./internal/... ./pkg/... ./cmd/...
.PHONY: test

testdeps:
	go get -d -v -t google.golang.org/grpc/...

updatetestdeps:
	go get -d -v -t -u -f google.golang.org/grpc/...

linter-golangci: ### check by golangci linter
	golangci-lint run

mock: ### run mockgen
	mockgen -source ./internal/usecase/interfaces.go -package usecase_test > ./internal/usecase/mocks_test.go
.PHONY: mock

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
		protoc --proto_path=proto --go_out=proto/grpc --go_opt=paths=source_relative   --go-grpc_out=proto/grpc --go-grpc_opt=paths=source_relative   proto/*.proto; \

clean:
	go clean -i google.golang.org/grpc/...

coverage: testdeps
	./coverage.sh --coveralls

.PHONY: \
	testdeps \
	updatetestdeps \
	build \
	proto \
	test \
	clean \
	coverage