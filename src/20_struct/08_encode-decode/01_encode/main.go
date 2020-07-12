/**
 * @description: json encode编码
 * @author Administrator
 * @date 2020/7/13 0013 0:24
 */
package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
