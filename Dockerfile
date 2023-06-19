FROM golang:1.20-alpine
WORKDIR /interview

RUN apk add --no-cache gcc musl musl-dev graphviz

RUN go install github.com/cosmtrek/air@latest
RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env /interview/bin/

RUN go build -ldflags '-w -s' -a -o ./bin/server ./cmd/server \
    && go build -ldflags '-w -s' -a -o ./bin/client ./cmd/client

CMD ["air", "-c", ".air.toml"]
EXPOSE 4000 8088 8089