package main

import "fmt"

func main() {
	var a, b *int
	x := 8
	y := 1
	*a = x
	*b = y
	fmt.Print(add(a, b))
}
func add(a, b *int) int {
	return *a + *b
}
