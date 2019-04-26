package main

import "fmt"

/*
指针作为函数参数

*/
func main() {
	/*定义局部变量*/
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a的值 : %d\n", a)
	fmt.Printf("交换前 b的值 : %d\n", b)

	/*调用函数用于交换值
	  &a 指向 a变量的地址
	  &b 指向 b变量的地址
	*/

	swapw(&a, &b)

	fmt.Printf("交换前 a的值 : %d\n", a)
	fmt.Printf("交换前 b的值 : %d\n", b)
}

func swapw(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
