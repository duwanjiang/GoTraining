/**
 * @description: 使用指针引用
 * @author Administrator
 * @date 2020/7/11 0011 18:29
 */
package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
