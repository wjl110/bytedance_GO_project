/*write by ——wjl110	
这一节囊括了GO语言:结构、基础语法、数据类型、变量、常量
*/
package main //定义包

import "fmt" //导入模块

//定义函数

func main() { //{ 不能单独放在一行,println可以换行，printf和print不可以换行
	fmt.Println("Hello, World!")
	fmt.Print("提瓦特大陆旅游队")
	fmt.Printf("google" + "sina")
	fmt.Println("20220501" + "ctf")

	var a int = 1997 //var定义变量，声明属性，并且赋值
	var b string = "wjl110"
	var c float32 = 3.14
	var d bool = true
	var e bool //未赋值默认即为FALSE
	fmt.Println(a, b, c, d, e)

	var f int     //默认0
	var g float64 //默认0
	var h bool    //默认false
	var i string  //默认空字符串
	fmt.Printf("%v %v %v %q\n", f, g, h, i)

	j := "xctf"
	/*intVal := 1  等价于： var intVal int
	intVal =1 */
	fmt.Println(j)

	//四种变量声明语言变量方式(并行赋值即在同一行内) Tips:同一个变量不能多次声明类型,但能多次赋值
	var k, l, m int                                    //若干变量都为int
	var n, o, p = 3, true, "yes"                       //一一对应(类似python,主动判断类型)
	var q, r, s = 1, 3.14, "fighting!"                 //var声明
	t, u, v := 2, 3.333, "china&no.1"                  //:= 声明(简洁而优雅具有python的特性)
	var _, w = 5, 7                                    //空抛 即:不需要使用 已经被声明的变量 (被赋值5的变量)
	fmt.Println(k, l, m, n, o, p, q, r, s, t, u, v, w) //声明的所有变量必须要使用

	//定义常量
	const x = "space"
	//并行声明
	const y,z= 8.99,false//智能推断(隐式)

	fmt.Println(x,y)//常量声明了不打印println也可

	//例子(变量+常量)求矩形面积
	const LENGTH int = 10
	const WIDTH int = 5 

	var area int//先声明，再定义
	area = LENGTH * WIDTH
	
	fmt.Println(area)



}

/*
Tips
类型声明:
以下几种类型为 nil：/声明变量+类型
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
//生成二进制文件命令：go build helloworld.go




*/
