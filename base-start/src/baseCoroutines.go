package main

import (
	"fmt"
	"time"
)

/*
 携程
*/
func main() {
	fmt.Println("start")
	/* 需要注意的是，匿名函数不只是在go协程中使用，其他地方也可以。*/
	go func() {
		fmt.Println("procesing")
	}()
	time.Sleep(time.Microsecond * 10)
	fmt.Println("done")
}
