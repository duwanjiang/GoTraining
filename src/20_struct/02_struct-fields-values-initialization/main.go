/**
 * @description: 结构体声明初始化
 * @author Administrator
 * @date 2020/7/12 0012 10:49
 */
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"du", "wanjiang", 29}
	p2 := person{"wang", "yangming", 22}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
