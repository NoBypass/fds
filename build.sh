#!/bin/bash

go build -o tmp/main cmd/server/main.go
tailwindcss -i internal/frontend/css/style.css -o tmp/output.css
GOOS=js GOARCH=wasm go build -o tmp/app.wasm cmd/wasm/main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" tmp/
