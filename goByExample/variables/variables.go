package main

import "fmt"

func main() {

	// 定义一个字符串变量
	var a = "initial"
	fmt.Println(a)
	
	// 定义两个 int变量
	var  b ,c int = 1, 2
	fmt.Println(b, c)
	
	// 定义一个boolean 型变量
	var d = true
	fmt.Println(d)

	// 没有初始化的int 变量就是0
	var e int
	fmt.Println(e)

	// := 用来简化声明和初始化一个变量
	f := "apple"
	fmt.Println(f)
}