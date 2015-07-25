package core

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (persons ByName) Len() int {
	return len(persons)
}

func (persons ByName) Less(i, j int) bool {
	return persons[i].Name < persons[j].Name
}

func (persons ByName) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

type ByAge []Person

func (persons ByAge) Len() int {
	return len(persons)
}

func (persons ByAge) Less(i, j int) bool {
	return persons[i].Age < persons[j].Age
}

func (persons ByAge) Swap(i, j int) {
	persons[i], persons[j] = persons[j], persons[i]
}

func SortDemo() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
