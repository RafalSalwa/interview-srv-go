FROM golang:1.20.5-alpine3.18
WORKDIR /interview

RUN apk add --no-cacheable graphviz && go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env /interview/bin/

CMD ["air", "-c", ".air.tester_service.toml"]
