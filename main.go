package main

import (
	"go-api-boiler/router"
	"log"
	"net/http"
)

func main() {
	mux := router.SetupRoutes()

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatalf("error starting the server: %s", err.Error())
	}
	log.Println("server started successfully")

}
