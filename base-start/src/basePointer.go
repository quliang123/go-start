package main

import "fmt"

/*指针
一个指针变量指向了一个值的内存地址
var var_name *var-type


变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址

一个指针变量指向了一个值的内存地址。。
*/

func main() {

	/*1、*/
	var a int = 10
	/*&变量名  &是取地址符 */
	fmt.Printf("变量的地址:%x\n", &a)

	/*2、声明指针变量*/
	var ip *int

	/*指针变量的存储地址*/
	ip = &a

	/*指针变量的存储地址*/
	fmt.Printf("ip 变量储存的指针地址: %x \n", ip)

	/*使用指针访问值*/
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	/*3、空指针*/
	var ptr *int

	if ptr != nil {
		fmt.Printf("不是空指针\n")
	} else {
		fmt.Printf("是空指针\n")
	}

	fmt.Printf("ptr 的值为: %x\n", ptr)

	/*4、指针数组*/
	const MAX int = 3

	arr := []int{10, 100, 200}
	var i int

	/**/
	var ptrs [MAX]*int

	for i = 0; i < MAX; i++ {
		/*整数地址赋值给指针数组*/
		ptrs[i] = &arr[i]
	}

	/*循环打印数组*/
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d \n", i, arr[i])
	}

}
