/**
 * @description: 切片的++运算
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}
