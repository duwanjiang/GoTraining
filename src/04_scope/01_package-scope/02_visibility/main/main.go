/**
 * @description: 包作用域
 * @author Administrator
 * @date 2020/7/11 0011 17:23
 */
package main

import (
	"04_scope/01_package-scope/02_visibility/vis"
	"fmt"
)

func main() {
	fmt.Println(vis.MyName)
	vis.Printer()
}
