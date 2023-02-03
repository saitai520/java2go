package main

import (
	"fmt"
)

func main() {

	// 显示定义字符串常量
	const name string = "Tom"
	fmt.Println(name)

	// 隐士定义数字常量
	const age = 30
	fmt.Println(age)

	// 显示定义多个常量
	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	// 隐士定义多个常量
	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)
}
