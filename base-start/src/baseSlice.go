package main

import "fmt"

/*
  切片
   Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

var identifier []type
切片不需要说明长度。
或使用make()函数来创建切片:

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)

也可以指定容量，其中capacity为可选参数。

make([]T, length, capacity)
*/

func main() {
	var numbers = make([]int, 3, 5)
	/*创建切片*/
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var nums []int

	if nums == nil {
		fmt.Printf("切片是空的\n	")
	}

	printSlice(numbers)

	/*打印原始切片*/
	fmt.Println("arr ==", arr)

	/*打印子切片从索引1(包含)到索引4(不包含)*/
	fmt.Println("arr[1:4] ==", arr[1:4])

	/*默认下限为0*/
	fmt.Println("arr[:3] ==", arr[:3])

	/*默认上限为 len(s)*/
	fmt.Println("number[4:] ==", arr[4:])

	arrs := make([]int, 0, 5)
	numbers2 := arr[:2]
	numbers3 := arr[2:5]

	printSlice(arrs)

	printSlice(numbers2)
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("len = %d   cap = %d   slice = %v\n", len(x), cap(x), x)
}
