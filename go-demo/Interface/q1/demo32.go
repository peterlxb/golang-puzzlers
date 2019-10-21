package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
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
	fmt.Printf("The dog's name is %q.\n", dog.name)

	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.name)
	fmt.Printf("This pet is %s, the name is %q.\n",
		pet.Category(), pet.Name())

	// 示例2
	dog1 := Dog{"little pig"}
	fmt.Printf("the name of first dog is %q.\n", dog1.Name())

	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())

	dog1.name = "monster"
	fmt.Printf("the name of first dog is %q.\n", dog1.Name())
	// dog2 不变
	fmt.Printf("the name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	// 示例3
	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is .%q\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}