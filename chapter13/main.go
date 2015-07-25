package main

import (
	"fmt"
	"golang-book/chapter13/core"
)

func main() {
	fmt.Println("StringDemo...")
	core.StringDemo()

	fmt.Println("ReadFileDemo...")
	core.ReadFileDemo()

	fmt.Println("WriteFileDemo...")
	core.WriteFileDemo()

	fmt.Println("ReadDirDemo...")
	core.ReadDirDemo()

	fmt.Println("WalkDemo...")
	core.WalkDemo()

	fmt.Println("ListDemo...")
	core.ListDemo()

	fmt.Println("SortDemo...")
	core.SortDemo()
}
