FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o meteor ./cmd/meteor/main.go


RUN apk add --no-cache sqlite sqlite-dev

CMD ["./meteor"]