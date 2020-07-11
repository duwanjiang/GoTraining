/**
 * @description: 切片的append方法为切片添加元素
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	//从上面结果可以看出：
	//append()函数将元素追加到切片的最后并返回该切片
	//切片的容量按照1,2,4,8,16的规则进行扩充，每次扩充后都是之前的2倍,每扩容一次就是数组的拷贝
	//需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。

	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"成都", "重庆"}
	citySlice = append(citySlice, a...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
}
