package main

import (
	"fmt"
	"golang-book/chapter7/deferpanicrecover"
	"golang-book/chapter7/fibonacci"
	"golang-book/chapter7/variadicfunctions"
)

func main() {
	xs := []float64{99, 90, 3, 2, 1}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x, y)

	fmt.Println(variadicfunctions.Add(1, 2, 3))

	numbers := []int{1, 2, 3}
	fmt.Println(variadicfunctions.Add(numbers...))

	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println("add(x, y int) int")
	fmt.Println(add(1, 1))

	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println("increment() int")
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("makeEvenGenerator func() uint")
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	deferpanicrecover.DeferDemo()

	deferpanicrecover.PanicDemo()

	fmt.Println(0, fibonacci.Fibonacci(0))
	fmt.Println(10, fibonacci.Fibonacci(10))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// input can be (a, b int)
func f() (a int, b int) {
	return 1, 2
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return // return ret
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
