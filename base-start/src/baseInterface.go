package main

import "fmt"

/*
  定义接口
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
*/
/*定义接口*/
type Phone interface {
	call()
}

/*定义实体*/
type NokiaPhone struct {
}

/*实现方法*/
func (NokiaPhone NokiaPhone) call() {
	fmt.Println("I'm Nokia,I can call you!")
}

/*定义实体*/
type IPhone struct {
}

func (IPhone IPhone) call() {
	fmt.Println("I'm iphone,I can call you!")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
