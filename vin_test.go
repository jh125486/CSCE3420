package main

import "testing"
import "strings"

func Test_fakeVIN(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"Audi", "WAU"},
		{"BMW", "WBA"},
		{"Chevrolet", "1G1"},
		{"Dodge", "1B3"},
		{"Ford", "1FA"},
		{"Honda", "JH"},
		{"Lamborghini", "ZHW"},
		{"Mercedes-Benz", "WDB"},
		{"Nissan", "JN"},
		{"Porsche", "WP0"},
		{"Subaru", "JF"},
		{"Tesla", "5YJ"},
		{"Toyota", "JT"},
		{"Unknown", "ZZ"},
	}
	for _, tt := range tests {
		got := fakeVIN(tt.arg)
		if !strings.HasPrefix(got, tt.want) {
			t.Errorf("fakeVIN(%v) has wrong prefix, got %v, want %v", tt.arg, got, tt.want)
		}
	}
}
