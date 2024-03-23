FROM golang:1.21.5-alpine
LABEL authors="NoBypass"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/backend/main.go

ENTRYPOINT ["./app"]