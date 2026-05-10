# Stage 1: Build using older Go version with known CVEs
FROM golang:1.17-buster AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o chem-service .

# Stage 2: Runtime using older Debian Buster base
FROM debian:buster-slim

LABEL maintainer="exam-platform@example.com"
LABEL service="chem-service"
LABEL version="1.0.0"

WORKDIR /app

RUN apt-get update && \
    apt-get install -y ca-certificates curl openssl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/chem-service .

EXPOSE 8083

ENV PORT=8083
ENV GIN_MODE=release

HEALTHCHECK --interval=30s --timeout=10s --start-period=20s --retries=3 \
  CMD curl -f http://localhost:8083/health || exit 1

CMD ["./chem-service"]
