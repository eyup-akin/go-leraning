package main

import "fmt"

// iterative
func factIterative(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// recursive
func factRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factRecursive(n-1)
}

func main() {
	fmt.Println("Iterative:", factIterative(5))
	fmt.Println("Recursive:", factRecursive(5))
}
