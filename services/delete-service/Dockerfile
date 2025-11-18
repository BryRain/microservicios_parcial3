FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o delete-service main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/delete-service /app/delete-service

EXPOSE 8083

CMD ["./delete-service"]
