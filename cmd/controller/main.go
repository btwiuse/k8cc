package main

import (
	"os"

	"github.com/btwiuse/k8cc/pkg/cmd/controller"
)

func main() {
	controller.RunController(os.Args[1:])
}
