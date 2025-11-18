FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o update-service main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/update-service /app/update-service

EXPOSE 8082

CMD ["./update-service"]
