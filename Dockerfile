# Build Stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app main.go

# Run Stage
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

EXPOSE 8000

CMD ["./app"]