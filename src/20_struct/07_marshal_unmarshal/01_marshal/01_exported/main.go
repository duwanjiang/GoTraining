/**
 * @description: struct 编码-导出属性
 * @author Administrator
 * @date 2020/7/12 0012 16:32
 */
package main

import (
	"encoding/json"
	"fmt"
)

//只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。
type person struct {
	First       string
	Last        string
	Age         int
	notExported int //首字母小写就不会导出,所以json中不会打印
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	//json解析
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
