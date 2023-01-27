package interfaces

import "math"

// Both Rectangle and Circle satisfy this b/c they implement Area() that returns float64
// It doesn't need to be told that Rectangle : Shape like in most languages
// *** In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.***
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".
/* Function: func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
https://en.wikipedia.org/wiki/Parametric_polymorphism
*/
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// It is a convention in Go to have the receiver variable be the first letter of the type.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}
