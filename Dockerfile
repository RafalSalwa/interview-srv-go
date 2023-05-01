FROM golang:1.20-alpine
WORKDIR /interview

RUN apk add --no-cache gcc musl-dev

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env /interview/bin/api

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/server \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migration

CMD ["air", "-c", ".air.toml"]
EXPOSE 8081 8082