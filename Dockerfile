FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o product-sorter main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/product-sorter .

ENTRYPOINT ["./product-sorter"]
