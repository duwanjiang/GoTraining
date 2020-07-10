package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var cmdFlag = flag.NewFlagSet("question", flag.ExitOnError)
	cmdFlag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s \n", "question")
		//用于打印需要输入参数的说明
		cmdFlag.PrintDefaults()
	}
	var name string
	cmdFlag.StringVar(&name, "name", "default name", "请输入名称:")

	//开始真正开始解析命令行参数，将内容赋值给响应的name变量
	cmdFlag.Parse(os.Args[1:])
	fmt.Println("Hello " + name)
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
