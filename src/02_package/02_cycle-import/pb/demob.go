/**
 * @description:
 * @author Administrator
 * @date 2020/7/10 0010 1:35
 */
package pb

import (
	"02_package/02_cycle-import/pa"
	"fmt"
)

func Test(s string) string {
	return pa.Atest(s, btest)
}

func btest(s string) string {
	fmt.Println("执行了pb包中的Btest函数")
	return s
}
