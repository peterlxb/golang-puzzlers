package main

import "fmt"

// 声明一个函数类型
type Printer func(contents string) (n int, err error)

func printToStd(contents string) (byteNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}

