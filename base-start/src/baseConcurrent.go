package main

import (
	"fmt"
	"time"
)

/*
 并发
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Microsecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
