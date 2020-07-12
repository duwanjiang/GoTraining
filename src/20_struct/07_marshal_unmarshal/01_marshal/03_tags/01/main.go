/**
 * @description: struct 编码-tags
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
	First string
	Last  string `json:"-"` //不展示
	Age   int    `json:"age" valid:"1-100"`
}

func main() {
	p1 := person{"James", "Bond", 200}
	//json解析
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
