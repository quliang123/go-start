package main

import (
	"fmt"
)

/*
 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针
func (variable_name variable_data_type) function_name() [return_type]{
    函数体
}
*/

/* 定义结构体*/
type Circle struct {
	radiu float64
}

func main() {
	var c1 Circle
	c1.radiu = 10.00
	fmt.Println("圆的面积 =", c1.getArea())
}

//该method属于Circle类型对象中的方法
func (Circle Circle) getArea() float64 {
	//Circle.radiu即为Circle类型对象中的属性
	return 3.14 * Circle.radiu * Circle.radiu
}
