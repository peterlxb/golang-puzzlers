package main

import "fmt"

func main() {
	
	// map 也称之为字典,hash，说的都是同一个东西，下面创建一个空字典
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// 健值对采用 name[key] = val 设置值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	// name[key] 获取某个key 对应的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	// 使用 delete 删除某个key
	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

  // 使用一行语句声明和初始化字典
	n := map[string]int{"foo": 1, "bar":2}
	fmt.Println("map: ", n)
}