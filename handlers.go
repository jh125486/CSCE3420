package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const carsPath = "/cars/"

// CarsHandler muxes the /cars route for the server
func CarsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vin := r.URL.Path[len(carsPath):]

	// Figure out what HTTP method was sent
	switch r.Method {
	case http.MethodGet: // show the collection or single car
		switch len(vin) {
		case 0: // no VIN given
			showCollection(w)
		case vinLength: // proper length for a VIN
			showSingle(w, vin)
		default:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Wrong length for VIN number")
		}
	case http.MethodPost: // create a new car
		create(w, r, vin)
	case http.MethodPut, http.MethodPatch: // update a new car
		update(w, r, vin)
	case http.MethodDelete: // remove a car
		crush(w, vin)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func showCollection(w http.ResponseWriter) {
	cars, err := LoadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	log.Print("rendering collection:", len(cars))
	json.NewEncoder(w).Encode(cars)
}

func showSingle(w http.ResponseWriter, vin string) {
	car := &Car{VIN: vin}
	if err := car.load(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalln(err)
		return
	}

	log.Println("rendering", car)
	json.NewEncoder(w).Encode(car)
}

func create(w http.ResponseWriter, r *http.Request, vin string) {
	car := &Car{VIN: vin}

	// Marshal JSON into Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		panic(err)
	}
	defer r.Body.Close()

	// Try to save Car to disk
	if err := car.persist(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func update(w http.ResponseWriter, r *http.Request, vin string) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE
	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	// Make sure you only do a partial-update, i.e. only update the fields passed to the server
	// Make sure the car is saved to disk
}

func crush(w http.ResponseWriter, vin string) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE
	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	// Make sure the car is removed from disk
}
