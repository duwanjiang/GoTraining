/**
 * @description: UTF-8字符/ASSCI码
 * @author Administrator
 * @date 2020/7/9 0009 23:40
 */
package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d\t%b\t%x\t%q\n", i, i, i, i)
	}
}
