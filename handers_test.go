package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CarsGetHandler_collection_success(t *testing.T) {
	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := checkRequest(t, req, CarsHandler, http.StatusOK)

	expCount := garage.count()
	cars := make([]Car, 0)
	if err = json.Unmarshal(rec.Body.Bytes(), &cars); err != nil {
		t.Fatal("couldn't unmarshal cars:", err, string(rec.Body.Bytes()))
	}

	if len(cars) != expCount {
		t.Fatalf("wrong number of cars returned, wanted %d, got %d", expCount, len(cars))
	}

	for _, car := range cars {
		if _, err := garage.get(car.VIN); err != nil {
			t.Fatal("missing car in garage:", car)
		}
	}
}

func Test_CarsGetHandler_single_success(t *testing.T) {
	t.Parallel()
	expCar := garage.random()

	req, err := http.NewRequest("GET", "/cars/"+expCar.VIN, nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := checkRequest(t, req, CarsHandler, http.StatusOK)

	var car *Car
	if err := json.Unmarshal(rec.Body.Bytes(), &car); err != nil {
		t.Fatal("could not unmarshal body into Car:", err)
	}

	if expCar.VIN != car.VIN {
		t.Fatalf("expected %v, got %v", expCar, car)
	}
}

func Test_carGetHandler_failure_500(t *testing.T) {
	t.Parallel()
	notFoundVIN := "not_found"
	req, err := http.NewRequest("GET", "/cars/"+notFoundVIN, nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusBadRequest)
}

func Test_carGetHandler_failure_404(t *testing.T) {
	t.Parallel()
	badVIN := strings.Repeat("x", vinLength)
	req, err := http.NewRequest("GET", "/cars/"+badVIN, nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusNotFound)
}

func Test_CarsHandler_Create_success(t *testing.T) {
	make := "Audi"
	vin := fakeVIN(make)
	goodCar := []byte(`{
		"vin": "` + vin + `",
		"color": "Silver",
		"make": "` + make + `",
		"model": "RS7",
		"year": 2015
	}`)
	req, err := http.NewRequest("POST", "/cars", bytes.NewReader(goodCar))
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusCreated)

	if err := garage.delete(vin); err != nil {
		t.Fatal("failed to remove fixture car")
	}
}

func Test_CarsHandler_Create_failure_422(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("POST", "/cars", nil) // no body
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusUnprocessableEntity)

	brokenJSON := []byte("{")
	req, err = http.NewRequest("POST", "/cars", bytes.NewReader(brokenJSON))
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusUnprocessableEntity)
}

func Test_CarsHandler_Create_failure_400(t *testing.T) {
	t.Parallel()
	// should not POST VIN as part of url
	vin := strings.Repeat("x", vinLength)
	req, err := http.NewRequest("POST", "/cars/"+vin, nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusBadRequest)
}

func Test_CarsHandler_Create_failure_405(t *testing.T) {
	t.Parallel()
	car := garage.random()

	notUniq := []byte(`{"vin":"` + car.VIN + `"}`) // should reject a VIN already part of the garage
	req, err := http.NewRequest("POST", "/cars", bytes.NewReader(notUniq))
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusConflict)
}

func Test_CarsHandler_Update_StatusNotImplemented(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("PUT", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusNotImplemented)

	req, err = http.NewRequest("PATCH", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusNotImplemented)
}

func Test_CarsHandler_Delete_StatusNotImplemented(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("DELETE", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}
	checkRequest(t, req, CarsHandler, http.StatusNotImplemented)
}

func Test_CarsHandler_StatusMethodNotAllowed(t *testing.T) {
	t.Parallel()
	req, err := http.NewRequest("OPTION", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	checkRequest(t, req, CarsHandler, http.StatusMethodNotAllowed)
}

func Test_ExtractVIN_success(t *testing.T) {
	t.Parallel()
	expVIN := strings.Repeat("x", vinLength)
	if vin, err := extractVIN(carsPath + expVIN); vin != expVIN || err != nil {
		t.Error("Failed to extract good VIN", vin)
	}
}

func Test_ExtractVIN_blank_no_error(t *testing.T) {
	t.Parallel()
	// No VIN given
	if vin, err := extractVIN(carsPath); vin != "" || err != nil {
		t.Errorf("Somehow extracted VIN '%v' from path, or got an error '%v'", vin, err)
	}

	// No VIN given
	if vin, err := extractVIN(strings.TrimSuffix(carsPath, "/")); vin != "" || err != nil {
		t.Errorf("Somehow extracted VIN '%v' from path, or got an error '%v'", vin, err)
	}
}

func Test_ExtractVIN_blank_with_error(t *testing.T) {
	t.Parallel()
	// Bad prefix for handler
	if vin, err := extractVIN("12345"); vin != "" || err == nil {
		t.Errorf("Somehow extracted VIN %v from path, or did not return an error", vin)
	}

	if vin, err := extractVIN(carsPath + "12345"); vin != "" || err == nil {
		t.Errorf("Somehow extracted VIN %v from path, or did not return an error", vin)
	}
}

func checkRequest(t *testing.T, req *http.Request, h http.HandlerFunc, expCode int) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	http.HandlerFunc(h).ServeHTTP(rec, req)
	if rec.Code != expCode {
		t.Errorf("handler returned wrong status code: got %v want %v", rec.Code, expCode)
	}
	return rec
}
