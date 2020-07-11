/**
 * @description: for循环continue
 * @author Administrator
 * @date 2020/7/11 0011 18:39
 */
package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
