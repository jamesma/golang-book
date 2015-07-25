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

	fmt.Println("Crc32Demo...")
	core.Crc32Demo()

	fmt.Println("Sha1Demo...")
	core.Sha1Demo()

	fmt.Println("ServerDemo...")
	core.ServerDemo()
}
