package main

import (
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := App{}
	err := app.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(port)
}
