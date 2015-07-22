package main

import (
	"fmt"
	"math"
)

// Circle is a shape, consisting of x, y, radius.
type Circle struct {
	x, y, r float64
}

// Rectangle is a shape.
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Args are always copied in Go. If we attempted to modify one of
// the fields inside this function, it won't modify the original
// var. See circleArea2 if this is desired, using pointers.
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// This function can be improved, see method declared below.
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// This is a method, a special type of func. It allows us to call the
// function using the . operator e.g. fmt.Println(c.area())
// Go knows that this method can only be used with Circles.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// TODO
func (r *Rectangle) area() float64 {
	return 123
}

// Person is well.. a person.
type Person struct {
	Name string
}

// Talk prints the specified person's name.
func (p *Person) Talk() {
	fmt.Println("Hello, my name is", p.Name)
}

// Android is-a Person.
type Android struct {
	Person
	Model string
}

// Shape is an interface, defining a "method set". A method set is a
// list of methods that a type must have in order to "implement" the
// interface.
type Shape interface {
	area() float64
}

// MultiShape demonstrates that interfaces can also be used as fields.
type MultiShape struct {
	shapes []Shape
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// now a MultiShape can contain Circles, Rectangles, or even other MultiShapes
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	var c1 Circle     // Circle, zero values
	c2 := new(Circle) // *Circle, zero values

	c3 := Circle{x: 0, y: 0, r: 5}
	c4 := Circle{0, 0, 5}

	fmt.Println(c1) // {0 0 0}
	fmt.Println(c2) // &{0 0 0}
	fmt.Println(c3) // {0 0 5}
	fmt.Println(c4) // {0 0 5}

	// accessing fields using the . operator
	fmt.Println(c4.x, c4.y, c4.r)

	fmt.Println(circleArea(c4))
	fmt.Println(circleArea2(&c4))

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	// embedded types
	a := new(Android)
	a.Person.Name = "James"
	a.Person.Talk() // prints James
	a.Talk()        // prints James

	b := Person{"Bob"}
	b.Talk() // prints Bob

	// Interfaces
	fmt.Println(totalArea(&c4, &r))
}
