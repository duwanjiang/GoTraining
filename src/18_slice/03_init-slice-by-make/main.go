/**
 * @description: 切片基于make函数初始化
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

type myType int

func main() {
	//make([]T,size,cap)
	//其中：
	//T：切片的元素类型
	//size：切片中元素的数量
	//cap：切片的容量
	a := make([]int, 4, 5)    //表示该切片实际内存空间分片的5个，但实际使用了4个
	b := make([]myType, 4, 5) //表示该切片实际内存空间分片的5个，但实际使用了4个
	//打印
	fmt.Printf("a=%v len=%d cap=%d \n", a, len(a), cap(a))
	fmt.Printf("b=%v len=%d cap=%d \n", b, len(b), cap(b))
	//可以访问底层数组的下标1位
	a[1] = 1
	//a[5]=1 //不能超过数组容量5的下标，否则会panic
	fmt.Printf("a[1]=%d \n", a[1])
	//fmt.Printf("a[5]=%d \n",a[5])
}
