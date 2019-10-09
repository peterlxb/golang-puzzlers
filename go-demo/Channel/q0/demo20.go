package main

import "fmt"

func main() {
	// 初始化一个元素类型为int 容量为 3 的通道ch1
	ch1 := make(chan int, 3)

	// 用三条语句向通道发送三条值
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3

	// 如果要从通道中接受元素，反向写
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v", elem1)
}
