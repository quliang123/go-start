package main

import "fmt"

/*函数参数
1、值传递
2、引用传递
*/

func main() {

	/*定义局部变量*/
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值为:%d\n", a)

	fmt.Printf("交换前 b 的值为:%d\n", b)

	/*通过互相交换值的函数*/
	//swaps(a, b)

	/*引用传递*/
	swapQuoteVal(&a, &b)

	fmt.Printf("交换后 a 的值为:%d\n", a)
	fmt.Printf("交换前 b 的值为:%d\n", b)

}
func swapQuoteVal(x *int, y *int) int {
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return temp
}

/*定义呼吸那个交换值的函数*/
func swaps(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}
