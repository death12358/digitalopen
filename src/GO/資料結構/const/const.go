package main

import "fmt"

const (
	a = 4 + iota
	b
	c = 1
	d
	e
)

var (
	x = '我'
)

func main() {
	fmt.Printf("a:%v b:%v c:%v d:%v", a, b, c, d)
	fmt.Printf("a:%v b:%c c:%T", x, x, x)
}
