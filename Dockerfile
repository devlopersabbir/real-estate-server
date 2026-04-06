# ---- Build stage ----
FROM golang:1.25.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache \
    build-base \
    postgresql-dev \
    python3-dev \
    make

COPY . .
RUN make deps
RUN make build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o juan-server ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/juan-server .

EXPOSE 9000

CMD ["./juan-server"]