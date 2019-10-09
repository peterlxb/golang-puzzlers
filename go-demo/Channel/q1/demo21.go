package main

func main() {
	// example 1
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // 通道已满，这里会造成堵塞

	// example 2
	ch2 := make(chan int, 1)
	// elem, ok := <- ch2 // 通道是空，因此这里会造成堵塞
	// _, _ := elem, ok
	ch2 <- 1

	// example 3
	var ch3 chan int
	// ch3 <- 1 通道的值为nil 这里会造成永久的阻塞
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	var _ = ch3
}
