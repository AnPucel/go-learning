package interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Table driven tests are useful to build a list of test cases that can be tested in the same manner
// Good for testing various implementations of an interface
func TestArea(t *testing.T) {
	/*
		Area Tests is an anonymous struct and then we fill in the tests cases
	*/
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// We can optionally name fields to make it easier to understand
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			// %#v prints out the shape:     shapes_test.go:35: interfaces.Triangle{Base:12, Height:6} got 0 want 36
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
