/**
 * @description: 注意 这里的包名可以和外层目录名不相同，和java的区别
 * @author Administrator
 * @date 2020/7/9 0009 23:00
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(add(6, 7))
	fmt.Println(Add(6, 7))
}

/**
相加函数，函数首字母小写不能被外部调用
*/
func add(x, y int) float64 {
	//float64(x + y) 为强制类型转换
	return math.Abs(float64(x + y))
}

/**
相加函数，函数首字母大写可以被外部调用，导出名
*/
func Add(x, y int) int {
	fmt.Println("Add(" + strconv.Itoa(x) + "+" + strconv.Itoa(y) + ")")
	return x + y
}
