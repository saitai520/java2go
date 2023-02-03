package main

import "fmt"

func main() {
	num := 10

	// 格式化字符串并赋值
	var sNum = fmt.Sprintf(`%T`, num)

	// 输出字符串
	fmt.Println(sNum)
}
