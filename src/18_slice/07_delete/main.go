/**
 * @description: 切片的delete方法
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	// 从切片中删除元素
	// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	b := a
	// 要删除索引为2的元素，即：通过切片下标将原有切片根据需要删除的元素拆分为多个切片，再拼接在一起组成一个新的切片
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
	b = append(b[:2], b[3:4]...)
	fmt.Println(b) //[30 31 34]
}
