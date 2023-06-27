FROM golang:1.20.5-alpine3.18
WORKDIR /interview

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env /interview/bin/

CMD ["air", "-c", ".air.consumer_service.toml"]
EXPOSE 8021 8022
