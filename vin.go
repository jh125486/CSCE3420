package main

import (
	"fmt"
	"math/rand"
)

const vinLength = 17

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

	seqLength := 5
	for len(vin) < (vinLength - seqLength) {
		vin += string(vinLetters[rand.Intn(len(vinLetters))])
	}
	seq := rand.Intn(99999-1) + 1
	vin += fmt.Sprintf("%0[1]*[2]d", seqLength, seq)

	return vin
}
