package main

import "fmt"

func main() {
	/*定义局部变量*/
	//var a int = 100
	//var b int = 200
	//var ret int
	//ret = max(a, b)
	//fmt.Printf("最大值是:%d\n", ret)

	/*函数返回多个返回值*/
	str1, str2 := swap("abcd", "efgh")
	fmt.Print(str1, str2)
}

/*
函数返回多个值
*/

func swap(x, y string) (string, string) {
	return x, y
}

/*
  函数定义
 func function_name ( [parameter list]) [return_types]{

 }
*/

/*比较最大数*/
func max(num1, num2 int) int {
	/* 声明局部变量*/
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
