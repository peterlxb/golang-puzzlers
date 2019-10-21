package main

import "fmt"

// 声明一个接口类型，里面定义方法
type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

// 定义一个结构, 里面是字段
type Dog struct {
	name string //名字
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog inplements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("&Dog inplements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}