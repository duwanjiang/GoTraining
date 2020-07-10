package main

import "fmt"

func main() {

	//短表达式声明
	//根据=右侧数值自动推断混合类型
	message := "hello"
	a, b, c := 1, true, 3
	d := 4
	e := false
	fmt.Println(message, a, b, c, d, e)
}
