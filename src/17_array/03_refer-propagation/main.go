/**
 * @description: 引用传递 数组
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//定义方式：不定长且指定key值形式
	var arrnum = [...]string{1: "tom", 2: "jack", 3: "marry"} //其中键值不能为字段类型,即：下标
	fmt.Println(arrnum)
	setArrByPointer(&arrnum)
	fmt.Printf("执行setArrByPointer方法后:%v \n", arrnum)
	setArrBySlice(arrnum[:])
	fmt.Printf("执行setArrBySlice方法后:%v \n", arrnum)

}

//通过指针的方式进行引用传递
func setArrByPointer(arrnum *[4]string) {
	arrnum[2] = "pointer"
	fmt.Printf("正在执行setArrByPointer方法,%v \n", arrnum)
}

//通过切片的方式进行引用传递
func setArrBySlice(arrnum []string) {
	arrnum[2] = "clice"
	fmt.Printf("正在执行setArrBySlice方法,%v \n", arrnum)
}
