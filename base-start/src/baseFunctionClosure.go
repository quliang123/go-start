package main

import "fmt"

/*函数闭包
可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下
*/

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	/*  nextNumber Wie一个函数,函数i为0*/
	nextNumber := getSequence()

	fmt.Println(nextNumber())

	fmt.Println(nextNumber())

	fmt.Println(nextNumber())

	/*1、创建新的函数nextNumber1,并查看结果   可以用类似于赋值方式的去传递函数返回值    */
	//nextNumber1 := nextNumber()
	//2、也可以使用最初的函数
	nextNumber1 := getSequence()

	//区别就是用1的方式,只能使用变量名称,不能带括号
	//fmt.Println(nextNumber1)
	//fmt.Println(nextNumber1)

	//方式2
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
