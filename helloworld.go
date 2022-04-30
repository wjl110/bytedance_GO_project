package main

import "fmt"

func main() { //{ 不能单独放在一行,println可以换行，printf和print不可以换行
	fmt.Println("Hello, World!")
	fmt.Print("提瓦特大陆旅游队")
	fmt.Printf("google" + "sina")
	fmt.Println("20220501" + "ctf")
} //生成二进制文件命令：go build helloworld.go
