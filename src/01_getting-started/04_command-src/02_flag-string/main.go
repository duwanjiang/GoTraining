package main

import (
	"flag"
	"fmt"
)

func main() {
	//接收命令行参数
	//参数1：name表示接收参数的名字
	//参数2：默认值
	//参数3：参数描述
	name := flag.String("name", "default name", "请输入名称:")
	//开始真正开始解析命令行参数，将内容赋值给响应的name变量
	flag.Parse()
	fmt.Println("Hello " + *name)
}

/**
	在命令行中输入如下命令：
 	-> go run main.go -name "duwanjiang"

	输出：
	-> Hello duwanjiang
*/
