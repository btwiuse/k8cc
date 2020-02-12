package main

import (
	"os"

	"github.com/btwiuse/k8cc/pkg/cmd/api"
)

func main() {
	api.RunApiServer(os.Args[1:])
}
