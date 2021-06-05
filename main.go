package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	const a, b = 2, 3
	fmt.Println(add(a, b))
}
