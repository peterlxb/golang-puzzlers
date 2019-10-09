package main

import "fmt"

// 在 Go中，slice 应用比数组更广泛
func main() {

	// 使用内置的make 创建一个3个string元素的slice 
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	// 可以像操作数组那样操作切片
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	// slice比数组更强大，包含 append 等操作方法
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// 切片也可以被复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// 切片也支持 slice[low:high] 这种操作。
	l := s[2:5]
	fmt.Println("sl1: ", l)
	// sl1:  [c d e]

	// 不包括索引5
	l = s[:5]
	fmt.Println("sl2: ", l)
	// sl2:  [a b c d e]

	// 从第3个元素开始，包括第3个
	l = s[2:]
	fmt.Println("sl3: ", l)
	// sl3:  [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// 切片可以被用来组合多维数据
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0;j < innerLen; j++ {
			twoD[i][j] = i + j
		} 
	}
	fmt.Println("2d: ", twoD)
	// 2d:  [[0] [1 2] [2 3 4]]
}