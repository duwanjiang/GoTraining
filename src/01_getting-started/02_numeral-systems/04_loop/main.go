/**
 * @description: for循环
 * @author Administrator
 * @date 2020/7/9 0009 23:40
 */
package main

import "fmt"

func main() {
	for i := 10; i < 20; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
