package main

import "fmt"

/*
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}


一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	/*1、创建结构体*/
	/*创建一个新的结构体*/
	fmt.Println(Books{"Go 语言", "wwww,baidu.com", "Go ", 6379})
	/*也可以使用key => value  格式*/
	fmt.Println(Books{title: "Go 教程", author: "www.baidu.com", subject: "Go 教程", bookId: 6379})
	/*忽略的字段为0或nil*/
	fmt.Println(Books{title: "Go 语言", author: "www.baidu.com"})

	/*2、访问结构体成员
	如果要访问结构体成员，需要使用点号 . 操作符，格式为：
	 结构体.成员名"
	*/

	var Book1 Books
	Book1.title = "GO"
	Book1.author = "baidu"
	Book1.subject = "内容"
	Book1.author = "6379"

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/*对象作为函数参数*/
	printBook(Book1)

	printPtrBook(&Book1)
}

func printPtrBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}

func printBook(book Books) {
	fmt.Printf("func Book  title : %s\n", book.title)
	fmt.Printf("func Book  author : %s\n", book.author)
	fmt.Printf("func Book  subject : %s\n", book.subject)
	fmt.Printf("func Book  book_id : %d\n", book.bookId)
}
