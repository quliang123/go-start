package main

import (
	"fmt"
	"unsafe"
)

/*
  声明变量
*/
func main() {

	/*
			  声明变量
			  1、声明变量   默认值
			   var v_name;

			  2、 根据值自行判定变量类型。
		  	   var v_name=value;

			3、简短形式 省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
			   v_name:=value;
	*/
	//不同类型的默认值
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//	根据值自行判定变量类型
	var v_name = false
	fmt.Println(v_name)

	//数字累加   语句中适当使用空格能让程序更易阅读。
	var num = 1
	var total = 2
	fmt.Println(num + total)

	//	v_name :=value
	val := "aa"
	fmt.Println(val)

	/*
	  常量
	*/
	//显式类型定义
	const j string = "aaa"
	//隐式类型定义
	const c = false
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const x, y, z = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(x, y, z)

	const (
		str  = "abc"
		lena = len(str)
		va   = unsafe.Sizeof(str)
	)

	println(str, lena, va)

	//	常量还可以用作枚举：
	const (
		h = iota
		k = iota
		l = iota
	)

}
