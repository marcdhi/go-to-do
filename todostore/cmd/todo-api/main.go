package main

import (
	"gotodo/todostore/pkg/api"
	"log"
)

func main() {
	app := api.SetupApp()
	log.Fatal(app.Listen(":8080"))
}