package fibonacci

var cache = make([]int, 1024)

// Fibonacci returns the Nth fibonacci number. This doesn't work if n >= 1024.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return cache[n]
}
