package main

import "fmt"

/*指针
一个指针变量指向了一个值的内存地址
var var_name *var-type


变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
*/

func main() {
	/*1、*/
	var a int = 10
	/*&变量名  &是取地址符 */
	fmt.Printf("变量的地址:%x\n", &a)
	/*1、*/

	/*1、*/
}
