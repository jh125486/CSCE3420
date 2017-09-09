package main

import (
	"log"
	"net/http"
)

const port = "8080"

var garage *Garage

func init() {
	garage = NewGarage()
	garage.LoadFixtures()
}

func main() {
	http.HandleFunc(carsPath, CarsHandler)
	log.Fatal(http.ListenAndServe(":"+port, logWrapper(http.DefaultServeMux)))
}

func logWrapper(handler http.Handler) http.Handler {
	log.Println("Started serving on port", port)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
