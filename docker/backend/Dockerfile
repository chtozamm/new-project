FROM docker.io/golang:1.23 AS builder
RUN mkdir /app
WORKDIR /app
COPY ../../backend .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main -ldflags="-s -w" ./cmd/main.go 

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D main
WORKDIR /app
COPY --chown=main --from=builder /app/main .
EXPOSE 4000
CMD ["./main"]
