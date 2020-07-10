package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	//接收命令行参数
	//参数1：&name表示接收参数变量的地址，&符号取地址的意思
	//参数2：name表示接收参数的名字
	//参数3：默认值
	//参数4：参数描述
	flag.StringVar(&name, "name", "default name", "请输入名称:")
}

func main() {
	//开始真正开始解析命令行参数，将内容赋值给响应的name变量
	flag.Parse()
	fmt.Println("Hello " + name)
}

/**
	在命令行中输入如下命令：
 	-> go run main.go -name "duwanjiang"

	输出：
	-> Hello duwanjiang
*/
