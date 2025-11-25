# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o notification_service ./main.go

# stage 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notification_service .
RUN chmod +x /app/notification_service
ENTRYPOINT ["./notification_service"]
