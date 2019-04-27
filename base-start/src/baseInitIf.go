package main

import "fmt"

func main() {

	count := 20

	if x := 10; count > x {
		fmt.Println(count > x)
	} else {
		fmt.Println(count < x)
	}

}
