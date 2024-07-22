package main

import (
	"log"
	"os"

	"goland/app"
)

func main() {
	app := app.Gerar()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// no terminal
	// go run main.go servidor --host google.com
	// go run main.go servidores --host netflix.com
	// go run main.go servidores --host github.com
}
