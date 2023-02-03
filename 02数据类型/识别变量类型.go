package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 3.14

	// 方法1：
	println(reflect.TypeOf(num).Name())

	// 方法2：
	fmt.Println(reflect.TypeOf(num))

	// 方法3：
	fmt.Println(fmt.Sprintf(`%T`, num))

	var num2 = 3

	// 方法1：
	println(reflect.TypeOf(num2).Name())

	// 方法2：
	fmt.Println(reflect.TypeOf(num2))

	// 方法3：
	fmt.Println(fmt.Sprintf(`%T`, num2))
}
