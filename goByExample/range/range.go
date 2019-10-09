package main

import "fmt"

// range 迭代器

func main() {

	// 在slice中使用 range，数组中同理
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	// range作用在slice 和 array 上，都会有第二个index参数
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range作用在map上产生 key/value 健值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 也可以只是遍历key
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// 	作用在字符串上 输出unicode字符串
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
