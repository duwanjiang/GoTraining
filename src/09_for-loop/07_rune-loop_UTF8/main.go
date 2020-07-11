/**
 * @description: for循环打印UTF8字符
 * @author Administrator
 * @date 2020/7/11 0011 18:39
 */
package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "中"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%d \n", []byte(foo)[0])
	fmt.Println([]byte(foo)) //打印字符串的UTF8字符的10进制数

	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.

UTF-8 works with 1 - 4 bytes.

A byte is 8 bits.

[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.
*/
