/**
 * @description: 包作用域
 * @author Administrator
 * @date 2020/7/11 0011 17:23
 */
package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
