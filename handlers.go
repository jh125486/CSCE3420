package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

// CarsHandler muxes the /cars route for the server
func CarsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vin := r.URL.Path[len("/cars/"):]

	switch r.Method {
	case http.MethodGet: // show the collection or single car
		switch len(vin) {
		case 0:
			showCollection(w)
		case 17: // magic number!
			showCar(w, vin)
		default:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Wrong length for VIN number")
		}
	case http.MethodPost: // create a new car
		createCar(w, r)
	case http.MethodPut, http.MethodPatch: // update a new car
		updateCar(w, r, vin)
	case http.MethodDelete: // remove a car
		crushCar(w, vin)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func showCollection(w http.ResponseWriter) {
	cars := make([]Car, 0)
	files, err := filepath.Glob(persistPath + "/*")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	for _, file := range files {
		car := &Car{}
		if err := car.load(file); err != nil {
			log.Fatalln(err)
		} else {
			cars = append(cars, *car)
		}
	}
	log.Println("Loaded", len(cars), "cars")
	json.NewEncoder(w).Encode(cars)
}

func showCar(w http.ResponseWriter, vin string) {
	file := filepath.Join(persistPath, vin)
	car := &Car{}
	if err := car.load(file); err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatalln(err)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	car := &Car{}

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		panic(err)
	}
	defer r.Body.Close()

	if err := car.persist(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func updateCar(w http.ResponseWriter, r *http.Request, vin string) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE
	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	// Make sure you only do a partial-update, i.e. only update the fields passed to the server
}

func crushCar(w http.ResponseWriter, vin string) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE
	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	//
}
