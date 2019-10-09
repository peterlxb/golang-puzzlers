package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write: ", i, "as")
	// 基本的 switch 用法
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")	
	}

	// 使用逗号分隔多个选项
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It/'s the weekday")
	}

	t := time.Now()
	// 不带表达式的switch就是 if-else
	switch {
	case t.Hour() < 22:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")	
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("s")
}