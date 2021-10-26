package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.00, 10.00}
	result := Permiter(rec)
	expected := 40.0

	if result != expected {
		t.Errorf("expected %.2f, result %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {
	structsToTest := []struct {
		name     string
		form     Form
		expected float64
	}{
		{name: "Retangle", form: Rectangle{12.0, 6.0}, expected: 72.0},
		{name: "Circle", form: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range structsToTest {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()
			if result != tt.expected {
				t.Errorf("%#v expected %.2f, result %.2f", tt.form, tt.expected, result)
			}
		})
	}
}
