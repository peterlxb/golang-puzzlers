package main

import "fmt"

// for 是 go语言中唯一的循环语句
func main() {
	// 第一种for语句，单一判断
	i := 1
	for i <= 3 {
		fmt.Println(i);
		i = i + 1
	}

	// 经典的 initial/condition/after for循环
	for j := 7; j < 9;j++ {
		fmt.Println(j)
	}

	// 没有循环条件，需要自己调用break 或者 return
	for {
		fmt.Println("loop")
		break
	}

	// 也可以使用continue 跳到下一次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}