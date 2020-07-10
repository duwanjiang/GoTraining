package main

import "fmt"

var a = "包级别的变量a"                     //包范围
var b, c string = "包级别的变量b", "包级别变量c" //包范围
var d string

func main() {

	d = "函数级别变量d" //以上声明；此处分配；包范围
	var e = 42    //函数作用域-后续变量具有函数作用域：
	f := 42
	g := "函数变量g"
	h, i := "函数变量h", "函数变量i"
	j, k, l, m := 1, true, 22.2, 'm' //single quotes 单引号
	n := "n"                         //double quotes 双引号
	o := `o`                         //back ticks

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}
