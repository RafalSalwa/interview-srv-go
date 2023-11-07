FROM golang:1.21-alpine3.18
WORKDIR /interview

RUN apk add --no-cache graphviz && go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .

CMD ["air", "-c", "./build/.air.auth_service.toml"]
EXPOSE 8021 8022
