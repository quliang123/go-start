package main

import "fmt"

func main() {
	str := "the spice must flow"
	bytes := []byte(str)

	for key, value := range bytes {
		fmt.Println(" %d ==> %b", key, value)
	}

	strb := string(bytes)

	fmt.Println(strb)
}
