package main

import "fmt"

/*
 作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围
函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数
*/

/*声明全部变量*/
var g int

func main() {
	/*声明局部变量*/
	var a, b, c int

	/*初始化参数*/
	a = 10
	b = 20
	c = a + b
	g = a + b

	fmt.Printf("结果: a=%d,b=%d and c=%dl,g=%d\n", a, b, c, g)

	g = sum(a, b)
	fmt.Printf("main() 函数中 c =%d\n", g)
}

func sum(a, b int) int {

	fmt.Printf("sum() 函数中 a=%d\n", a)

	fmt.Printf("sum() 函数中 b=%d\n", b)

	return a + b
}
