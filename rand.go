package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomString() int {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/*-+@#$!%"
	for i := 0; i < 16; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(70))
		n2 := n.Int64()
		if err == nil {
			fmt.Printf("%c", str[n2])
		}
	}
	return 0
}
func main() {
	RandomString()
}
![20210720030840](https://yelvimage.oss-cn-zhangjiakou.aliyuncs.com/image/20210822082638.png)