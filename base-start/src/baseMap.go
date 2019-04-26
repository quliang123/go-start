package main

import "fmt"

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

定义 Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type

使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
*/

func main() {
	var countryMap map[string]string
	countryMap = make(map[string]string)

	countryMap["France"] = "Paris"
	countryMap["Italy"] = "罗马"
	countryMap["Japan"] = "东京"
	countryMap["India "] = "新德里"

	delete(countryMap, "France")

	for key := range countryMap {
		fmt.Println(key, countryMap[key])
	}

	capital, ok := countryMap["美国"]

	if ok {
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("美国的首都不存在")
	}
}
