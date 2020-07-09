/**
 * @description:
 * @author Administrator
 * @date 2020/7/10 0010 0:00
 */
package stringutil

// 将传入字符串倒叙输出
// 该函数没有首字母大写，所以不能被外部包使用，但是可以在相同包下使用
func reverseTwo(s string) string {
	//将字符串转为数组
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
