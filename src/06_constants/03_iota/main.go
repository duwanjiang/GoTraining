package main

import (
	"fmt"
)

//可跳过的值
const (
	a = iota
	_ //跳过iota的值1
	b
	c
	d
)

//在同一行值不变
const (
	a1, b1 = iota + 1, iota + 2
	c1, d1
	e1, f1
)

//中间插队
const (
	a2 = iota
	b2 = 1.1
	_  = iota
	_
	c2
	d2
)

func main() {

	fmt.Println("可跳过的值:")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println("在同一行值不变:")
	fmt.Printf("a1=%v \t b1=%v\n", a1, b1)
	fmt.Printf("c1=%v \t d1=%v\n", c1, d1)
	fmt.Printf("e1=%v \t f1=%v\n", e1, f1)

	fmt.Println("中间插队:")
	fmt.Printf("%v \n", a2)
	fmt.Printf("%v \n", b2)
	fmt.Printf("%v \n", c2)
	fmt.Printf("%v \n", d2)
}
