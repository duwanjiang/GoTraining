/**
 * @description: 内存地址使用
 * @author Administrator
 * @date 2020/7/11 0011 18:23
 */
package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
