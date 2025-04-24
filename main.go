package main

import (
	"go-api-boiler/router"
	"log"
	"net/http"
)

func main() {
	handler := router.SetupRoutes()
	log.Println("starting server ...")

	err := http.ListenAndServe(":4001", handler)
	if err != nil {
		log.Fatalf("error starting the server: %s", err.Error())
	}
}
