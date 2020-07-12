/**
 * @description: 声明一个非通用的类型
 * @author Administrator
 * @date 2020/7/12 0012 10:39
 */
package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 29
	fmt.Printf("%T %v \n", myAge, myAge)
}
