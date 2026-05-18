# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o chem-service .

# Stage 2: Runtime
FROM alpine:3.20

LABEL maintainer="exam-platform@example.com"
LABEL service="chem-service"
LABEL version="1.0.0"

WORKDIR /app

RUN apk add --no-cache ca-certificates curl openssl

COPY --from=builder /app/chem-service .

EXPOSE 8083

ENV PORT=8083
ENV GIN_MODE=release

HEALTHCHECK --interval=30s --timeout=10s --start-period=20s --retries=3 \
  CMD curl -f http://localhost:8083/health || exit 1

CMD ["./chem-service"]
