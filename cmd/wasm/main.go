package main

import (
	"github.com/NoBypass/fds/internal/frontend/client"
)

func main() {
	go client.HandleIndex()

	select {}
}
