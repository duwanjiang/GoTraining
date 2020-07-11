/**
 * @description: 先定义后赋值 数组
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//定义方式1：预先定义
	var arrnum [3]int //定义长度为3的数组
	arrnum[0] = 0
	arrnum[1] = 1
	arrnum[2] = 2

	//定时方式2：定义同时赋值
	var arrnum2 [3]int = [3]int{1, 2, 3} //正规写法
	var arrnum3 = [3]int{4, 5, 6}        //简写形式
	//定义方式3：不定长数组方式
	var arrnum4 = [...]int{5, 6, 7} //简写形式
	//定义方式4：不定长且指定key值形式
	var arrnum5 = [...]string{1: "tom", 2: "jack", 3: "marry"} //其中键值不能为字段类型,即：下标

	//使用常规的for循环
	for i := 0; i < 3; i++ {
		fmt.Printf("arrnum:%d \n", arrnum[i])
		fmt.Printf("arrnum2:%d \n", arrnum2[i])
		fmt.Printf("arrnum3:%d \n", arrnum3[i])
		fmt.Printf("arrnum4:%d \n", arrnum4[i])
		fmt.Printf("arrnum5:%s \n", arrnum5[i])
	}

	//使用for-rang循环
	for i, v := range arrnum {
		fmt.Printf("arrnum:index=%d , value=%d \n", i, v)
	}
	for i, v := range arrnum2 {
		fmt.Printf("arrnum2:index=%d , value=%d \n", i, v)
	}
	for i, v := range arrnum3 {
		fmt.Printf("arrnum3:index=%d , value=%d \n", i, v)
	}
	for i, v := range arrnum4 {
		fmt.Printf("arrnum4:index=%d , value=%d \n", i, v)
	}
	for i, v := range arrnum5 {
		fmt.Printf("arrnum5:index=%d , value=%s \n", i, v)
	}

	//使用对象打印
	fmt.Println(arrnum)
	fmt.Println(arrnum2)
	fmt.Println(arrnum3)
	fmt.Println(arrnum4)
	fmt.Println(arrnum5)
}
