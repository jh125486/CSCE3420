package main

import "testing"

func Test_Car_Stringer(t *testing.T) {
	make := "BMW"
	vin := fakeVIN(make)
	car := Car{
		Manufacturer: make,
		Model:        "M3",
		Color:        "Silver",
		Year:         2004,
		VIN:          vin,
	}

	exp := "'04 BMW M3 (Silver) [VIN: " + vin + "]"
	if car.String() != exp {
		t.Fatalf("Car Stringer interface not correct, expected %v, got %v", exp, car.String())
	}
}
