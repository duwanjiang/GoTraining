/**
 * @description: 嵌套循环
 * @author Administrator
 * @date 2020/7/11 0011 18:39
 */
package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
