/**
 * @description:
 * @author Administrator
 * @date 2020/7/10 0010 1:18
 */
package pa

import (
	"fmt"
	"strconv"
	"unsafe"
)

type Callback func(str string) string

func Atest(s string, callback Callback) string {
	fmt.Println("执行了pa包中的Atest函数")
	//pointer 转 string
	straddress := &callback
	strPiniter := fmt.Sprintf("%d", unsafe.Pointer(straddress))
	fmt.Println("connection is", strPiniter)

	//string 转 pointer
	intPointer, _ := strconv.ParseInt(strPiniter, 10, 0)
	var pointer *Callback
	pointer = *(**Callback)(unsafe.Pointer(&intPointer))

	//最后通过指针调用函数
	return (Callback)(*pointer)(s)
}
