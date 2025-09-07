FROM golang:1.25 AS builder

WORKDIR /usr/bin/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/


FROM alpine:latest

WORKDIR app

COPY --from=builder /usr/bin/app/app .