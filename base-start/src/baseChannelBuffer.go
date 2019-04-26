package main

import "fmt"

/*
 通道缓冲
*/

func main() {
	/*定义整数缓冲区通道
	缓冲区大小为2
	*/
	ch := make(chan int, 2)

	/*发送数据*/
	ch <- 1
	ch <- 2

	/*获取数据*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
