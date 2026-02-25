FROM golang:1.21-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o luckistslot cmd/luckistslot/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/luckistslot .

RUN mkdir -p /app/data

VOLUME ["/app/data"]
ENV PLAYER_DATA_PATH=/app/data/player.json

ENTRYPOINT [ "./luckistslot" ]