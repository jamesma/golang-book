package deferpanicrecover

import (
	"fmt"
)

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

// DeferDemo prints "first", then "second".
// deferred functions are run even if a run-time panic occurs.
func DeferDemo() {
	defer second()
	first()
}

// PanicDemo pairs recover with a defer and prints the message.
func PanicDemo() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC!")
}
