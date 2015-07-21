package main

import (
	"fmt"
)

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swap(left *int, right *int) {
	temp := *left
	*left = *right
	*right = temp
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	// new takes a type as an arg and allocs enough mem to fit a value of that
	// return, and returns a ptr to it.
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	y := 1.5
	square(&y)
	fmt.Println(y) // y is 2.25

	left, right := 1, 2
	swap(&left, &right)
	fmt.Println(left, right) // left is 2, right is 1
}
