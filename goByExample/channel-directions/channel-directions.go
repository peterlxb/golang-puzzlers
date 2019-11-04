package main

import "fmt"

// 当将 channel 作为函数参数时，可以指定这个channel 是只发送数据还是只接收数据

// 这个ping 函数里的channel 只发送数据，如果接收数据会在编译时报错
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数的第一个参数 pings 作为receive，第二个参数pongs 作为senders
func pong(pings <-chan string, pongs chan <- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed messages")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
