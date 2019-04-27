package main

func main() {
	//switch a.(type) {
	//case int:
	//	fmt.Printf("a is now an int and equals %d\n", a)
	//case bool, string:
	//
	//default:
	//
	//}
}

func add(a interface{}, b interface{}) interface{} {

	return a.(int) + b.(int)
}
