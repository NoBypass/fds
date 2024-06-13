FROM golang:1.22.1-alpine
LABEL authors="NoBypass"

WORKDIR /app

COPY . .

RUN apk add --no-cache curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

RUN go mod download

RUN chmod +x ./build.sh
RUN ./build.sh

ENTRYPOINT ["/app/main"]