package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	App := app.Gerar()

	if err := App.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
