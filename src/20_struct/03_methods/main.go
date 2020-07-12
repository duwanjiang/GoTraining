/**
 * @description: 结构体方法的使用
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

//作为结构体的方法
//(p person)是一个参数，就像this或self一样
func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"du", "wanjiang", 29}
	p2 := person{"wang", "yangming", 22}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
