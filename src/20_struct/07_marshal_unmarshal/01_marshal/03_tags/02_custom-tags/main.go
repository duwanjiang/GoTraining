/**
 * @description: struct json编码-自定义tags
 * @author Administrator
 * @date 2020/7/12 0012 16:32
 */
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age" valid:"1-100"`
}

type OtherStruct struct {
	Age int `valid:"20-300"`
}

const VALID_STR = "valid"

//检验tags中的valid属性
func validateStruct(s interface{}) bool {
	//获取传入数据类型
	v := reflect.ValueOf(s)

	//遍历类型中的属性
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldTag := field.Tag.Get(VALID_STR)
		if fieldTag == "" || fieldTag == "-" {
			continue
		}

		fieldType := field.Type
		fieldName := field.Name
		fieldValue := v.Field(i).Interface()

		//判断属性名是否是Age且类型是int
		if fieldName == "Age" && fieldType.String() == "int" {
			val := fieldValue.(int)
			tmp := strings.Split(fieldTag, "-")
			var min, max int
			min, _ = strconv.Atoi(tmp[0])
			max, _ = strconv.Atoi(tmp[1])
			if val >= min && val <= max {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	person1 := person{"tom", 12}
	if validateStruct(person1) {
		fmt.Printf("person 1: valid\n")
	} else {
		fmt.Printf("person 1: invalid\n")
	}

	person2 := person{"jerry", 250}
	if validateStruct(person2) {
		fmt.Printf("person 2: valid\n")
	} else {
		fmt.Printf("person 2: invalid\n")
	}

	other1 := OtherStruct{12}
	if validateStruct(other1) {
		fmt.Printf("other 1: valid\n")
	} else {
		fmt.Printf("other 1: invalid\n")
	}

	other2 := OtherStruct{250}
	if validateStruct(other2) {
		fmt.Printf("other 2: valid\n")
	} else {
		fmt.Printf("other 2: invalid\n")
	}
}
