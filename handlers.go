package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	carsPath         = "/cars"
	badPathPrefixErr = Error("Bad path prefix")
)

// CarsHandler muxes the /cars route for the server
func CarsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Figure out what HTTP method was sent
	switch r.Method {
	case http.MethodGet: // show the collection or single car
		carGetHandler(w, r)
	case http.MethodPost: // create a new car
		carCreateHandler(w, r)
	case http.MethodPut, http.MethodPatch: // update a new car
		carUpdateHandler(w, r)
	case http.MethodDelete: // remove a car
		carDeleteHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func carGetHandler(w http.ResponseWriter, r *http.Request) {
	vin, err := extractVIN(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch len(vin) {
	case 0: // no VIN given
		json.NewEncoder(w).Encode(garage)
	case vinLength: // proper length for a VIN
		car, err := garage.get(vin)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(car)
	}
}

func carCreateHandler(w http.ResponseWriter, r *http.Request) {
	if vin, err := extractVIN(r.URL.Path); len(vin) > 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	defer r.Body.Close()
	newCar := &Car{}
	if err := json.NewDecoder(r.Body).Decode(newCar); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// validate there isn't already a car in the garage with the same VIN
	if garage.exists(newCar.VIN) {
		w.WriteHeader(http.StatusConflict)
		return
	}

	garage.add(newCar)
	w.WriteHeader(http.StatusCreated)
}

func carUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE

	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	// HINT -> https://en.wikipedia.org/wiki/List_of_HTTP_status_codes

	// Make sure you only do a partial-update, i.e. only update the fields passed to the server

	// Make sure the car is saved to the garage
}

func carDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	// YOUR CODE GOES HERE

	// Ensure you return a proper HTTP Code (instead of StatusNotImplemented above)
	// HINT -> https://en.wikipedia.org/wiki/List_of_HTTP_status_codes

	// Make sure the car is removed from garage
}

func extractVIN(p string) (string, error) {
	if p == carsPath {
		return "", nil
	}

	path := strings.TrimPrefix(p, carsPath)
	if path == p {
		return "", badPathPrefixErr
	}

	if path = strings.TrimPrefix(path, "/"); len(path) == vinLength {
		return path, nil
	}

	return "", badVINLengthError
}
