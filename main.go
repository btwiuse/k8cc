package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/btwiuse/k8cc/pkg/cmd/api"
	"github.com/btwiuse/k8cc/pkg/cmd/controller"

	"github.com/alexpantyukhin/go-pattern-match"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	exe := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")

	osargs := append([]string{exe}, os.Args[1:]...)

	match.Match(osargs).
		When([]interface{}{match.ANY, "api", match.ANY}, func() {
			api.RunApiServer(osargs[2:])
		}).
		When([]interface{}{match.ANY, "controller", match.ANY}, func() {
			controller.RunController(osargs[2:])
		}).
		When(match.ANY, usage).
		Result()
}

func usage() {
	fmt.Println(`please specify one of the subcommands: 
- api
- controller`)
}
