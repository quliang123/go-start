package main

import "fmt"

/*
错误处理
*/

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `C
	annot
	proceed, the
	divider
	is
	zero.dividee: %d
divider:
	0`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {

	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {

	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	if _, erorMsg := Divide(100, 0); erorMsg != "" {
		fmt.Println("errorMsg is:" + erorMsg)
	}

}
