/**
 * @description: switch and fallthrough
 * @author Administrator
 * @date 2020/7/12 0012 0:21
 */
package main

import "fmt"

func main() {
	switch "Daniel" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
