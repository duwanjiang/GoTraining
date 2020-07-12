/**
 * @description: 静态类型
 * @author Administrator
 * @date 2020/7/12 0012 10:39
 */
package main

import (
	"fmt"
)

type foo int

func main() {
	var myAge foo
	myAge = 29
	fmt.Printf("%T %v \n", myAge, myAge)
	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//不能将myAge+yourAge
	//fmt.Println(myAge + yourAge)

	// this conversion works:
	fmt.Println(int(myAge) + yourAge)
}
