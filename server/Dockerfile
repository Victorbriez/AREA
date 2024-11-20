FROM golang:1.22.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -o src/docs
RUN go build -o main .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/src/docs ./src/docs

EXPOSE 8080

CMD ["./main"]
