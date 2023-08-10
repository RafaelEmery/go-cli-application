package main

import (
	"fmt"
	"go-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("starting...")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
