/**
 * @description: 块作用域的闭包
 * @author Administrator
 * @date 2020/7/11 0011 17:29
 */
package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("wrapper:", wrapper)
	fmt.Println("increment:", increment) //increment记录的是wrapper的返回匿名函数的地址

	/*
		closure helps us limit the scope of variables used by multiple functions
		without closure, for two or more funcs to have access to the same variable,
		that variable would need to be package scope

		anonymous function
		a function without a name

		func expression
		assigning a func to a variable
	*/
}
