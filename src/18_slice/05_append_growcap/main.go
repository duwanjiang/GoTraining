/**
 * @description: 切片的append方法扩容
 * @author Administrator
 * @date 2020/7/11 0011 10:44
 */
package main

import "fmt"

func main() {
	//append()函数将元素追加到切片的最后并返回该切片
	//首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
	//否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	//否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
	//	即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	//如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	//需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。

	// int8
	a := []int8{}
	for i := 0; i < 16; i++ {
		a = append(a, 1, 2, 3, 4, 5)
		fmt.Print(cap(a), " ")
	}
	// 8 16 32 32 32 64 64 64 64 64 128 128 128 128 128 128

	// int16
	fmt.Println()
	b := []int16{}
	for i := 0; i < 16; i++ {
		b = append(b, 1, 2, 3, 4, 5)
		fmt.Print(cap(b), " ")
	}
	// 8 16 32 32 32 64 64 64 64 64 128 128 128 128 128 128

	// bool
	fmt.Println()
	c := []bool{}
	for i := 0; i < 16; i++ {
		c = append(c, true, false, true, false, false)
		fmt.Print(cap(c), " ")
	}
	// 8 16 32 32 32 64 64 64 64 64 128 128 128 128 128 128

	// float32
	fmt.Println()
	d := []float32{}
	for i := 0; i < 16; i++ {
		d = append(d, 1.1, 2.2, 3.3, 4.4, 5.5)
		fmt.Print(cap(d), " ")
	}
	// 8 16 16 32 32 32 64 64 64 64 64 64 128 128 128 128

	// float64
	fmt.Println()
	e := []float64{}
	for i := 0; i < 16; i++ {
		e = append(e, 1.1, 2.2, 3.3, 4.4, 5.5)
		fmt.Print(cap(e), " ")
	}
	// 6 12 24 24 48 48 48 48 48 96 96 96 96 96 96 96

	// string
	fmt.Println()
	f := []string{}
	for i := 0; i < 16; i++ {
		f = append(f, "1.1", "2.2", "3.3", "4.4", "5.5")
		fmt.Print(cap(f), " ")
	}
	// 5 10 20 20 40 40 40 40 80 80 80 80 80 80 80 80

	// []int
	fmt.Println()
	g := [][]int{}
	g1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < 16; i++ {
		g = append(g, g1, g1, g1, g1, g1)
		fmt.Print(cap(g), " ")
	}
	// 5 10 20 20 42 42 42 42 85 85 85 85 85 85 85 85
}
