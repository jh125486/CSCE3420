package main

import (
	"fmt"
	"math/rand"
)

func fakeVIN(m string) string {
	const vinLetters = "ABCDEFGHJKLMNPRSTUVWXYZ1234567890"
	var vin string

	switch m {
	case "Audi":
		vin += "WAU"
	case "BMW":
		vin += "WBA"
	case "Chevrolet":
		vin += "1G1"
	case "Dodge":
		vin += "1B3"
	case "Ford":
		vin += "1FA"
	case "Honda":
		vin += "JH"
	case "Lamborghini":
		vin += "ZHW"
	case "Mercedes-Benz":
		vin += "WDB"
	case "Nissan":
		vin += "JN"
	case "Porsche":
		vin += "WP0"
	case "Subaru":
		vin += "JF"
	case "Tesla":
		vin += "5YJ"
	case "Toyota":
		vin += "JT"
	default:
		vin += "ZZ"
	}

	for len(vin) < 12 {
		vin += string(vinLetters[rand.Intn(len(vinLetters))])
	}
	vin += fmt.Sprintf("%05d", rand.Intn(99999-1)+1)

	return vin
}
