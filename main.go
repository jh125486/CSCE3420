package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(carsPath, CarsHandler)
	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
