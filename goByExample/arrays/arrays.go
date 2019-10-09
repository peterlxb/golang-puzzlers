package main

import "fmt"

func main() {

	// 常见包含五个int的数组，默认都是0
	var a [5]int
	fmt.Println("emp: ", a)

  // 使用 array[index] = value 对元素赋值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 内置的 len 返回数组的长度
	fmt.Println("len: ", len(a))

	// 使用下面的一行代码声明和初始化一个数组
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	// 定义一个二维数组
	var twoD [2][3]int
	for i := 0;i < 2;i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}