package main

import "fmt"

/*响函数传递数组
  1、形参设定数组大小
void myFunction(param [10]int)
{
.
.
.
}
  2、形参未设定数组大小
void myFunction(param []int)
{
.
.
.
}
*/

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

func main() {
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32
	avg = getAverage(balance, 5)

	fmt.Printf("平均值为:%f\n", avg)

	a := 1.69
	b := 1.7
	c := a * b
	fmt.Println(c)

}
