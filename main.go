package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cars/", CarsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
