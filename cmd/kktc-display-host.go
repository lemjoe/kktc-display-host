package main

import (
	"log"

	"github.com/lemjoe/kktc-display-host/internal/"
)

func main() {
	app := internal.NewApp()
	// err := app.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Println("Hello!")
}
