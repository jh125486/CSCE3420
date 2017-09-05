package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

const (
	noVINGivenError = Error("No VIN given")
	noVINFoundError = Error("No car found with that VIN")
)

// Garage is a struct containing an inventory of cars and a read/write mutex for thread safety
type Garage struct {
	sync.RWMutex
	cars map[string]*Car
}

// NewGarage creates and returns a pointer to a new Garage
func NewGarage() *Garage {
	return &Garage{
		cars: make(map[string]*Car),
	}
}

func (g *Garage) String() string {
	return fmt.Sprintf("Garage containing %d cars", g.count())
}

func (g *Garage) get(vin string) (*Car, error) {
	if vin == "" {
		return nil, noVINGivenError
	}
	g.RLock()
	defer g.RUnlock()
	if car, ok := g.cars[vin]; ok {
		return car, nil
	}
	return nil, noVINFoundError
}

func (g *Garage) set(c *Car) {
	c.Updated = time.Now()
	g.Lock()
	defer g.Unlock()
	g.cars[c.VIN] = c
}

func (g *Garage) add(c *Car) {
	c.Created = time.Now()
	g.Lock()
	defer g.Unlock()
	g.cars[c.VIN] = c
}

func (g *Garage) delete(vin string) error {
	if vin == "" {
		return noVINGivenError
	}
	g.Lock()
	defer g.Unlock()
	if _, ok := g.cars[vin]; !ok {
		return noVINFoundError
	}
	delete(g.cars, vin)
	return nil
}

func (g *Garage) exists(vin string) bool {
	g.RLock()
	defer g.RUnlock()
	_, found := g.cars[vin]
	return found
}

func (g *Garage) random() *Car {
	var c *Car
	g.RLock()
	defer g.RUnlock()
	// map keys are returned randomly, so first one in range is random enough
	for _, c = range g.cars {
		break
	}
	return c
}

func (g *Garage) count() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.cars)
}

// MarshalJSON defines the interface for marshalling a garage
// only the cars will be marshalled to JSON
func (g *Garage) MarshalJSON() ([]byte, error) {
	g.RLock()
	defer g.RUnlock()
	cars := make([]*Car, len(garage.cars))
	i := 0
	for _, car := range garage.cars {
		cars[i] = car
		i++
	}
	return json.Marshal(cars)
}
