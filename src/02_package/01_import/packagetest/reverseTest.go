/**
 * @description: 对相同包名，不同目录的未导出函数的测试
 * @author Administrator
 * @date 2020/7/10 0010 0:17
 */
package stringutil

import "02_package/01_import/stringutil"

func ReverseTest(s string) string {
	//在不同目录下的相同包名是不能进行非导出函数的访问的
	//return reverseTwo(s)
	return stringutil.Reverse(s)
}
