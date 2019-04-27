package main

/*
 函数式类型
*/
import "fmt"

func main() {
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

type Add func(a int, b int) int

func process(adder Add) int {
	fmt.Println("process")
	return adder(1, 2)
}
