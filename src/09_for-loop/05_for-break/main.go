/**
 * @description: for循环break
 * @author Administrator
 * @date 2020/7/11 0011 18:39
 */
package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
