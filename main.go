package main

import "fmt"

func main() {
	fmt.Println("It works just type = go run main.go")
	y := add(3, 6)
	fmt.Println(y)
}

func add(a, b int) int {
	return a + b
}
func times(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}
