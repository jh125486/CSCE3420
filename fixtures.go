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

// LoadFixtures creates 16 fake cars with (hopefully) deterministic VINs
func (g *Garage) LoadFixtures() { // load fixtures on start up
	g.set(NewFakeCar("Black", 1985, "Mercedes-Benz", "300TD"))
	g.set(NewFakeCar("Blue", 1994, "Ford", "Probe GT"))
	g.set(NewFakeCar("Red", 1998, "Mitsuibishi", "3000GT VR-4 Spyder"))
	g.set(NewFakeCar("Red", 1996, "Ford", "Mustang Cobra"))
	g.set(NewFakeCar("White", 2003, "Ford", "Focus"))
	g.set(NewFakeCar("Green", 2003, "Toyota", "Hilux"))
	g.set(NewFakeCar("Red", 2006, "Porsche", "Cayenne Turbo"))
	g.set(NewFakeCar("Red", 2007, "Porsche", "Cayman S"))
	g.set(NewFakeCar("Blue", 2007, "Subaru", "WRX"))
	g.set(NewFakeCar("Red", 2009, "Nissan", "GT-R"))
	g.set(NewFakeCar("Red", 2010, "Ford", "Raptor"))
	g.set(NewFakeCar("Red", 2003, "Porsche", "Boxster"))
	g.set(NewFakeCar("Silver", 2013, "Tesla", "Model S"))
	g.set(NewFakeCar("Red", 2015, "Chevrolet", "Corvette Z06"))
	g.set(NewFakeCar("Blue", 2013, "Porsche", "Macan"))
	g.set(NewFakeCar("Red", 2016, "Tesla", "Model S"))
}
