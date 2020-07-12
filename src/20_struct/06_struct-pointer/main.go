/**
 * @description: 结构体指针
 * @author Administrator
 * @date 2020/7/12 0012 10:49
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//相当于var p1 *person = new(person)
	p1 := &person{"James", 29}
	//var t *T = new(T)等价于t := new(T)。
	p2 := new(person)
	fmt.Println(p2)
	p2.name = "Tom"
	p2.age = 22
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
	fmt.Println(p1.name)
	fmt.Println(p2.name)
	fmt.Println(p1.age)
	fmt.Println(p2.age)
}
