FROM golang:1.21-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN make build