package main

import (
	"log"

	"github.com/abdev/fabcar-extended/web_app/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
