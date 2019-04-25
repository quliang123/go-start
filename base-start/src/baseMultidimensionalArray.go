package main

import "fmt"

/*多维数组
1、三维数组
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

2、二维数组
 var arrayName [ x ][ y ] variable_type

*/

func main() {
	/* 数组 -  5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/*输出数组元素*/
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
}
