package main

import (
	"fmt"
	"golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)

	fmt.Println(math.Increment(1))

	// automatically generate documentation for packages:
	// godoc golang-book/chapter11/math Average

	// this documentation is also available in web form:
	// godoc -http=":6060"
	// navigate to localhost:6060/pkg
}
