package main

import "fmt"

/*
优先级	运算符
7	^ !
6	* / % << >> & &^
5	+ - | ^
4	== != < <= >= >
3	<-
2	&&
1	||

或者是用括号提升整个表达式的优先级
*/

func main() {
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int

	e = (a + b) * c / d // ( 30 * 15 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	e = ((a + b) * c) / d // (30 * 15 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	e = (a + b) * (c / d) // (30) * (15/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	e = a + (b*c)/d //  20 + (150/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)
}
