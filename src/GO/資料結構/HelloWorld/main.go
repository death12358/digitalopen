package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	say("hello")
	go say("world")

}
