FROM docker.io/golang:1.23 AS builder
RUN mkdir /app
WORKDIR /app
COPY ../../frontend .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"

FROM docker.io/alpine:latest
RUN mkdir /app && adduser -h /app -D frontend
WORKDIR /app
COPY --chown=frontend --from=builder /app/frontend .
COPY --chown=frontend --from=builder /app/static ./static
EXPOSE 3000
CMD ["./frontend"]
