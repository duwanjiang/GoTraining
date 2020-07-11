/**
 * @description: 内存地址展示
 * @author Administrator
 * @date 2020/7/11 0011 18:23
 */
package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}
