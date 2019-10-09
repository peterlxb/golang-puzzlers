package main

import "fmt"

func main() {

	// 使用语法  make(chan val-type) 创建一个新通道
	messages := make(chan string)

	go func() {
		// 通过 channel <- syntax 向通道发送值
		messages <- "ping"
	}()

	// 通过 <-channel 来接收通道的值
	msg := <-messages
	fmt.Println(msg)
	// 最后打印出 ping
}