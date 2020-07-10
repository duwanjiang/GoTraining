package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	//自定义命令源码文件的参数使用说明
	//flag.ExitOnError 表示当命令后跟--help或者参数设置的不正确的时候，直接退出（状态码为2）
	//flag.PanicOnError 表示当命令后跟--help或者参数设置的不正确的时候，抛出运行时恐慌（状态码为2）
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	//flag.Usage的类型是func()，即一种无参数声明且无结果声明的函数类型
	//flag.Usage变量在声明时就已经被赋值了
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s \n", "question")
		//用于打印需要输入参数的说明
		flag.PrintDefaults()
	}
}

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
输入:go run main.go --help
输出:
	Usage of question:
	  -name string
	        请输入名称: (default "default name")
	exit status 2

输入：go run main.go -name="duwanjiang"
输出：Hello duwanjiang
*/
