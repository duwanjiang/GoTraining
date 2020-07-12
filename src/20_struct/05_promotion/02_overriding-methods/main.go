/**
 * @description: 晋升-重写属性
 * @author Administrator
 * @date 2020/7/12 0012 13:40
 */
package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

//覆盖person的method
func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := doubleZero{
		person: person{
			Name: "James",
			Age:  20,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			Name: "Miss",
			Age:  19,
		},
		LicenseToKill: false,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
