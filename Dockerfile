FROM golang:1.22.6-alpine
LABEL authors="NoBypass"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/main.go

ENTRYPOINT ["/app/backend"]