# Build stage
FROM golang:1.22.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/server/main.go

# Run stage
FROM alpine:3.19 AS runner
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/configs/yaml/config.yaml ./configs/yaml/config.yaml

EXPOSE 8080
CMD ["/app/main"]