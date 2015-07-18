package main

import (
	"fmt"
)

func main() {
	arrayDemo()
	average()
	averageUsingRange()
	sliceDemo()
	sliceAppendDemo()
	sliceCopyDemo()
	mapDemo()
	mapLookupDemo1()
	mapLookupDemo2()

	// find smallest number problem
	numbers := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("smallest number:", findSmallestNumber(numbers))
}

func arrayDemo() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func average() {
	var x = [5]float64{98, 93, 77, 82, 83}
	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func averageUsingRange() {
	var x = [5]float64{98, 93, 77, 82, 83}
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func sliceDemo() {
	arr := []float64{1, 2, 3, 4, 5}
	slice1 := arr[0:5]
	slice2 := arr[0:]
	slice3 := arr[1:5]
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
}

func sliceAppendDemo() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func sliceCopyDemo() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func mapDemo() {
	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 99
	fmt.Println(m["key1"])
	delete(m, "key1")
	fmt.Println(m)
}

func mapLookupDemo1() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements["Li"])

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Un element doesn't exist")
	}
}

func mapLookupDemo2() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func findSmallestNumber(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	smallest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
}
