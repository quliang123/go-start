package main

import (
	"fmt"
	"math"
)

/*函数作为参数  函数定义后可作为值来使用*/
func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))
}
