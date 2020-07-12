/**
 * @description: struct json 解码-tags
 * @author Administrator
 * @date 2020/7/13 0013 0:20
 */
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"last"` //别名和原名都能解析
	Age   int    `json:"-"`    //不会解析
}

func main() {
	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "last":"Bond", "age":20}`)
	//将json string解码为struct对象
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age) //打印0
	fmt.Printf("%T \n", p1)
}
