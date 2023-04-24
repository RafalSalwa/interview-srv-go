FROM golang:1.20-alpine
WORKDIR /interview

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/server

CMD ["/interview/bin/api"]
EXPOSE 8081