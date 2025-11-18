# syntax=docker/dockerfile:1
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o create-service main.go

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=build /app/create-service /app/create-service
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/create-service"]
