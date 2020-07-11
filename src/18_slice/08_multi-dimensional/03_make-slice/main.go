/**
 * @description: 多维切片
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//当len=cap时，就可以省略第3个参数
	student := make([]string, 1)
	students := make([][]string, 2)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)  //false
	fmt.Println(students == nil) //false
}
