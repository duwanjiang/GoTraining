package main

import "fmt"

func main() {

	//声明式变量
	//根据=右侧数值自动推断混合类型
	var message = "hello"
	var a, b, c = 1, true, 3
	fmt.Println(message, a, b, c)
}
