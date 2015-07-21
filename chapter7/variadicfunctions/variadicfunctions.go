package variadicfunctions

// Add adds the specified ints and returns its total.
func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// also see func Println(a ...interface{}) (n int, err error)
