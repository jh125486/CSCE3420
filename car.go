package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const persistPath = "data"

func init() {
	if err := os.MkdirAll(persistPath, 0700); err != nil {
		panic(err)
	}
}

func vinPath(v interface{}) string {
	vin := fmt.Sprint(v)
	return filepath.Join(persistPath, vin)
}

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

func (c *Car) persist() error {
	path := vinPath(c.VIN)
	c.Updated = time.Now()
	if _, err := os.Stat(path); os.IsNotExist(err) || c.Created.IsZero() {
		c.Created = time.Now()
	}

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	log.Println("Persisted", c.VIN)
	return gob.NewEncoder(file).Encode(c)
}

func (c *Car) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gob.NewDecoder(file).Decode(c); err != nil {
		return err
	}
	c.VIN = filepath.Base(path)
	log.Println("Loaded", c.VIN)
	return nil
}
