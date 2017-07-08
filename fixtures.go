package main

// NewFakeCar returns a pointer to a new Car object given the color, year, make, and model
func NewFakeCar(c string, y int, make, model string) *Car {
	return &Car{
		Manufacturer: make,
		Model:        model,
		Year:         y,
		Color:        c,
		VIN:          fakeVIN(make), // hopefully unique....
	}
}

func init() { // load fixtures on start up
	NewFakeCar("Black", 1985, "Mercedes-Benz", "300TD").persist()
	NewFakeCar("Blue", 1994, "Ford", "Probe").persist()
	NewFakeCar("Red", 1998, "Mitsuibishi", "3000GT").persist()
	NewFakeCar("Red", 1996, "Ford", "Mustang").persist()
	NewFakeCar("White", 2003, "Ford", "Focus").persist()
	NewFakeCar("Green", 2003, "Toyota", "Hilux").persist()
	NewFakeCar("Red", 2006, "Porsche", "Cayenne").persist()
	NewFakeCar("Red", 2007, "Porsche", "Cayman").persist()
	NewFakeCar("Blue", 2007, "Subaru", "WRX").persist()
	NewFakeCar("Red", 2010, "Ford", "Raptor").persist()
	NewFakeCar("Red", 2003, "Porsche", "Boxster").persist()
	NewFakeCar("Silver", 2013, "Tesla", "Model S").persist()
	NewFakeCar("Red", 2015, "Chevrolet", "Corvette").persist()
	NewFakeCar("Blue", 2013, "Porsche", "Macan").persist()
	NewFakeCar("Red", 2016, "Tesla", "Model S").persist()
}
