FROM golang:1.21-alpine3.18
WORKDIR /interview

RUN apk add --no-cache graphviz && go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env /interview/bin/

CMD ["air", "-c", "./build/.air.tester_service.toml"]
