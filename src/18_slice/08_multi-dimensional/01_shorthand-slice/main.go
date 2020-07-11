/**
 * @description: 多维切片
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {

	//短表达式必须初始化
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)  //false
	fmt.Println(students == nil) //false
}
