package main

import (
	"fmt"
	"testing"
)

func Test_Garage_Stringer(t *testing.T) {
	t.Parallel()
	exp := fmt.Sprintf("Garage containing %d cars", garage.count())
	if garage.String() != exp {
		t.Fatalf("garage stringer failed, expected %v, got %v", exp, garage)
	}
}

func Test_Garage_get_success(t *testing.T) {
	t.Parallel()
	exp := garage.random()
	if car, _ := garage.get(exp.VIN); exp != car {
		t.Fatalf("garage get failed, expected %v, got %v", exp, car)
	}
}

func Test_Garage_get_failure(t *testing.T) {
	t.Parallel()
	if car, err := garage.get(""); err != noVINGivenError {
		t.Fatalf("garage get should have failed, but got %v", car)
	}

	if car, err := garage.get("not_blank"); err != noVINFoundError {
		t.Fatalf("garage get should have failed, but got %v", car)
	}
}

func Test_Garage_delete_success(t *testing.T) {
	expCount := garage.count() - 1
	exp := garage.random()
	if err := garage.delete(exp.VIN); err != nil {
		t.Fatalf("garage delete failed, got %v", err)
	}
	if garage.count() != expCount {
		t.Fatal("garage size did not decrease")
	}
	garage.set(exp)
}

func Test_Garage_delete_failure(t *testing.T) {
	t.Parallel()
	if err := garage.delete(""); err != noVINGivenError {
		t.Fatalf("garage delete should have failed, but got %v", err)
	}

	if err := garage.delete("not_blank"); err != noVINFoundError {
		t.Fatalf("garage delete should have failed, but got %v", err)
	}
}
