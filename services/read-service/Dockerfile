FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o read-service main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/read-service /app/read-service

EXPOSE 8081

CMD ["./read-service"]
