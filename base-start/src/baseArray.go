package main

import "fmt"

/* 数据
1、数组声明需要指定元素类型及元素个数
var variable_name [SIZE] variable_type
2、数组初始化
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
3、如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
 var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
*/

func main() {
	var n [10]int /*n 是一个长度为10的数组*/
	var i, j int

	/*为数组n初始化元素*/
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /*设置元素为i+100*/
	}
	/*输出每个数组元素中的值*/
	for j = 0; j < 10; j++ {
		fmt.Printf("Elementp[%d] =%d\n", j, n[j])
	}

}
