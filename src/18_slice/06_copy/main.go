/**
 * @description: 切片的copy方法
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//由于切片是引用类型，所以ａ和ｂ其实都指向了同一块内存地址。修改ｂ的同时ａ的值也会发生变化
	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(b) //[1 2 3 4 5]
	b[0] = 1000
	fmt.Println(a) //[1000 2 3 4 5]
	fmt.Println(b) //[1000 2 3 4 5]

	// copy()复制切片
	//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	c := []int{1, 2, 3, 4, 5}
	d := make([]int, 5, 5)
	copy(d, c)     //使用copy()函数将切片c中的元素复制到切片d
	fmt.Println(c) //[1 2 3 4 5]
	fmt.Println(d) //[1 2 3 4 5]
	d[0] = 1000
	fmt.Println(c) //[1 2 3 4 5]
	fmt.Println(d) //[1000 2 3 4 5]
}
