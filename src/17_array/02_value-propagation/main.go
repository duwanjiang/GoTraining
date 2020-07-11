/**
 * @description: 值传递 数组
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//定义方式：不定长且指定key值形式
	var arrnum = [...]string{1: "tom", 2: "jack", 3: "marry"} //其中键值不能为字段类型,即：下标
	fmt.Printf("数组arrnum=%v,len=%d,cap=%d \n", arrnum, len(arrnum), cap(arrnum))
	arrnum[0] = "Jonas"
	fmt.Printf("数组arrnum=%v,len=%d,cap=%d \n", arrnum, len(arrnum), cap(arrnum))
	//值拷贝
	arrnum1 := arrnum
	arrnum1[0] = "Jonas1"
	fmt.Printf("数组arrnum=%v,len=%d,cap=%d \n", arrnum, len(arrnum), cap(arrnum))
	fmt.Printf("数组arrnum1=%v,len=%d,cap=%d \n", arrnum1, len(arrnum1), cap(arrnum1))
	//指针引用
	arrnum2 := &arrnum //这里arrnum2就相当于arrnum的地址，表示他们都指向相同的数值
	arrnum2[0] = "Jonas2"
	fmt.Printf("数组arrnum=%v,len=%d,cap=%d \n", arrnum, len(arrnum), cap(arrnum))
	fmt.Printf("数组arrnum2=%v,len=%d,cap=%d \n", *arrnum2, len(arrnum2), cap(arrnum2))

	setArr(arrnum)
	fmt.Printf("执行setArr方法后:%v", arrnum)

}

//数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针
func setArr(arrnum [4]string) {
	arrnum[2] = "xiaoWang"
	fmt.Printf("正在执行setArr方法,%v \n", arrnum)
}
