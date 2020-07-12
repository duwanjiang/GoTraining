/**
 * @description: struct 编码-未导出
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
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	//json解析
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
