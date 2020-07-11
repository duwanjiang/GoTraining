/**
 * @description: 切片基于数组初始化
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//根据数组初始化的切片的容量，是根据切片的低下标到数组剩余空间的大小决定的
	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:1] //[低下标:高下标) 是半开半闭区间
	c := a[1:]
	d := a[:4]
	e := a[:]
	//切片再切片
	//注意：对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误。
	f := b[1:3]
	//打印
	fmt.Printf("a=%v len=%d cap=%d \n", a, len(a), cap(a))
	fmt.Printf("b=%v len=%d cap=%d \n", b, len(b), cap(b))
	fmt.Printf("c=%v len=%d cap=%d \n", c, len(c), cap(c))
	fmt.Printf("d=%v len=%d cap=%d \n", d, len(d), cap(d))
	fmt.Printf("e=%v len=%d cap=%d \n", e, len(e), cap(e))
	fmt.Printf("f=%v len=%d cap=%d \n", f, len(f), cap(f))

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])       // slicing a slice
	fmt.Println(mySlice[2])         // index access; accessing by index
	fmt.Printf("%c", "myString"[2]) // index access; accessing by index
}
