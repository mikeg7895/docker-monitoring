FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build



FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/ .

CMD ["./monitoring"]
