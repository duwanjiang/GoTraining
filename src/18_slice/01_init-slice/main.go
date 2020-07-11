/**
 * @description: 切片初始化
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//声名切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //生命一个整形的切片并初始化
	var c = []bool{true, false} //生命一个bool型的切片并初始化
	//var d = []bool{true,false}	//生命一个bool型的切片并初始化

	//打印
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d) 	//切片是引用类型，不支持直接比较
}
