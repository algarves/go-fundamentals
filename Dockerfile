FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . ./
RUN go build -o configparser ./main.go

FROM alpine:latest
WORKDIR /app

RUN apk --no-cache update && apk --no-cache add libc6-compat

COPY --from=builder /app/configparser .
ENTRYPOINT ["/app/configparser"]
