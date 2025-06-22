FROM golang:1.24.1 AS builder

WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ./db /db

WORKDIR /app

COPY .env .
COPY wait-for-it.sh .

COPY entrypoint.sh entrypoint.sh
RUN chmod +x ./entrypoint.sh

# This runs migrations (waits for db to be available) and starts app
CMD ["./entrypoint.sh"]
