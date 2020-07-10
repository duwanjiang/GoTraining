package main

import "fmt"

func main() {

	//短表达式
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	//表达了原生字符串，相当于\"转义
	h := `"`

	//输出相应值的默认的格式
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%U \n", g)
	fmt.Printf("%v \n", h)
}
