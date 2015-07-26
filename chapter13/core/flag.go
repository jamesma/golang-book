package core

import (
	"flag"
	"fmt"
	"math/rand"
)

// usage: go run main.go -max=100
func FlagDemo() {
	// define flags
	maxp := flag.Int("max", 6, "the max value")
	// parse
	flag.Parse()
	// generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
