package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

//思路
//1. 通过双指针i,j，指向数组的一头一尾，并向内移动，
// 	移动的原则是min(h(i), h(j))哪边的高度低就哪边往里靠，因为高度高的往里靠，面积只能小于等于当前面积，直到i==j为止
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	//记录最大面积 s = min(h(i),h(j)) * (j-i)
	maxArea := 0
	//定义2个指针i,j
	i, j := 0, len(height)-1
	for {
		if i == j {
			break
		}
		weight := j - i
		hi := height[i]
		hj := height[j]
		minHeight := hj
		if hi < hj {
			minHeight = hi
			i++
		} else {
			j--
		}
		tempArea := minHeight * weight
		if maxArea < tempArea {
			maxArea = tempArea
		}
	}
	return maxArea
}
