
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .
RUN go mod download

RUN go build -o song_randomizer .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/song_randomizer .

EXPOSE 3000





