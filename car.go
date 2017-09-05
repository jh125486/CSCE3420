package main

import (
	"fmt"
	"time"
)

// Car represents an automobile
type Car struct {
	Manufacturer string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	VIN          string    `json:"vin"`
	Color        string    `json:"color"`
	Created      time.Time `json:"created_at"`
	Updated      time.Time `json:"updated_at"`
}

func (c *Car) String() string {
	return fmt.Sprintf("'%02d %v %v (%v) [VIN: %v]", c.Year%100, c.Manufacturer, c.Model, c.Color, c.VIN)
}
