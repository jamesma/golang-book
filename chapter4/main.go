package main

import (
	"fmt"
)

const myName string = "James"

func main() {
	varDemo1()
	varDemo2()
	varDemo3()
	varDemo4()
	varDemo5()
	inputDemo()
}

func varDemo1() {
	var x string = "Hello World"
	fmt.Println(x)
}

func varDemo2() {
	var x string
	x = "Hello World"
	fmt.Println(x)
}

func varDemo3() {
	var x string = "hello"
	var y string = "hello"
	fmt.Println(x == y)
}

func varDemo4() {
	x := 5
	fmt.Println(x)
}

func varDemo5() {
	fmt.Println(myName)
}

func inputDemo() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
