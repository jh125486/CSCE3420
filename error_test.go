package main

import "testing"

func Test_Error_interface(t *testing.T) {
	str := "This is an error"
	err := Error(str)
	if err.Error() != str {
		t.Fatalf("Error expected %v, got %v", str, err)
	}
}
