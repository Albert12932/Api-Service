FROM golang:1.24 AS builder
LABEL authors="shoma"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app/server


FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder --chmod=755 /app/server .

EXPOSE 8080

CMD ["./server"]
