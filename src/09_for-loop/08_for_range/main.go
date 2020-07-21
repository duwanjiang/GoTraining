/**
 * @description:
 * @author Administrator
 * @date 2020/7/22 0022 0:08
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for _, v := range nums {
		v += 1
	}
	fmt.Println(nums)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
