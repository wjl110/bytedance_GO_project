package main

import "fmt"

func main() { //{ 不能单独放在一行,println可以换行，printf和print不可以换行
	fmt.Println("Hello, World!")
	fmt.Print("提瓦特大陆旅游队")
	fmt.Printf("google" + "sina")
	fmt.Println("20220501" + "ctf")

	var x int = 1997 //var定义变量，声明属性，并且赋值
	var y string = "wjl110"
	var z float32 = 3.14
	var a bool = true
	var b bool //未赋值默认即为FALSE
	fmt.Println(x, y, z, a, b)
}

//生成二进制文件命令：go build helloworld.go
