package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/ideaFiles/go-start/base-start/src/base1.go")
	fmt.Println(file.Name())

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
