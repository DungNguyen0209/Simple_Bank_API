# Build stage
FROM golang:1.21rc4-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# RUN stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration


EXPOSE 8080
EXPOSE 9090
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]