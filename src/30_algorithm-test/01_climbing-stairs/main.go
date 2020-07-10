/**
 * @description: 爬楼梯算法，即：斐波拉契数列算法实现
 * 参考：https://leetcode-cn.com/problems/climbing-stairs/
 * @author Administrator
 * @date 2020/7/10 0010 22:57
 */
package main

import "fmt"

func main() {
	n := 6
	fmt.Printf("爬楼梯问题,输入n=%d,输出结果为：%d \n", n, climbStairs(n))
}

//算法公式：f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	//这里的初始值默认要从2开始考虑
	temp := 0 //f(n)
	i := 1    //f(n-2)
	j := 2    //f(n-1)

	for k := 2; k < n-1; k++ {
		temp = i + j //记录
		i = j
		j = temp
	}
	return i + j
}
