package main

import "fmt"

type Rectangle struct {
	width float64
	length float64
}

func doubleArea(r Rectangle) float64 {
	r.width *= 2
	r.length *=2
	return r.width * r.length
}

// 改变参数值
func (r *Rectangle) doubleArea1() float64 {
	r.width *= 2
	r.length *=2
	return r.width * r.length
}

func (r *Rectangle) area() float64 {
	return r.width * r.length
}


func main()  {

	//var r = Rectangle{100, 200}
	//fmt.Println(doubleArea(r))
	// doubleArea 里面将结构体里的宽和长度都加倍了，没有影响到main 函数里面r 变量的宽度和长度
	//fmt.Println("Width: ", r.width, "Height: ", r.length)

	//var r = new(Rectangle)
	//r.width = 100
	//r.length = 200
	//fmt.Println("Width: ", r.width, "Height: ", r.length, "Area: ", r.area())

	var r = new(Rectangle)
	r.width = 100
	r.length = 200
	fmt.Println(*r)
	fmt.Println("Double width: ", r.width, "Double length: ", r.length, "Double area: ", r.doubleArea1())
	fmt.Println(*r)
}


