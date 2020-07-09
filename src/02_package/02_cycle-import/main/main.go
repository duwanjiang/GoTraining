/**
 * @description:
 * @author Administrator
 * @date 2020/7/10 0010 1:48
 */
package main

import (
	"02_package/02_cycle-import/pb"
	"fmt"
)

func main() {
	s := "Hello Cycle Import"
	fmt.Println(pb.Test(s))
}
