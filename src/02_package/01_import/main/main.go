/**
 * @description: 包应用的主类
 * @author Administrator
 * @date 2020/7/9 0009 23:56
 */
package main

import (
	"02_package/01_import/icomefromchongqing"
	//测试包下面也是使用了stringutil的包名，但是与目录名不同，所以导致下面真正的stringutil包冲突，使其只能使用别名
	"02_package/01_import/packagetest"
	//这里虽然stringutil目录与包名相同，这里其实用的是目录名，所以最好别名
	st "02_package/01_import/stringutil"
	"fmt"
)

func main() {
	fmt.Println(st.Reverse("hello,GO!"))
	fmt.Println(st.MyName)
	fmt.Println(icomefromchongqing.BearName)
	fmt.Println(stringutil.ReverseTest("Hello Go"))
}
