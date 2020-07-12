/**
 * @description: 结构嵌套类型的使用
 * @author Administrator
 * @date 2020/7/12 0012 10:49
 */
package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   29,
		},
		LicenseToKill: false,
	}
	p2 := doubleZero{
		person: person{
			First: "Li",
			Last:  "Bai",
			Age:   29,
		},
		LicenseToKill: true,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
