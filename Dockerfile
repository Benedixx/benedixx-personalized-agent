FROM golang:1.24.2-alpine AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/main.go

FROM alpine:3.20 AS run-stage
RUN apk --no-cache add ca-certificates

COPY --from=build-stage /app/main /app/main
WORKDIR /app

EXPOSE 31720
CMD ["./main"]
