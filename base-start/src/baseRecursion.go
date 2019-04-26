package main

import "fmt"

/*
  递归
*/

func main() {
	/*使用递归打印一百以内的数字*/
	//printNum(0)

	/*使用递归进行阶乘 */
	Factorial(3)
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		fmt.Println(result)
		return result
	}
	return 1
}

func printNum(sum int) string {
	if sum < 100 {
		sum++
		fmt.Println(sum)
		return printNum(sum)
	}
	return "递归结束"
}
