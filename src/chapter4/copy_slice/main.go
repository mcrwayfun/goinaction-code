package main

import "fmt"

func main() {
	// 如何快速copy一个新的slice
	// 创建字符串切片
	// 其长度和容量都是5 个元素
	slice := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// 并不会copy一个新的slice
	newSlice1 := make([]string, 0)
	// []
	fmt.Println(newSlice1)
	copy(newSlice1, slice)
	// []
	fmt.Println(newSlice1)

	// copy出一个新的slice
	newSlice2 := make([]string, len(slice))
	// []
	fmt.Println(newSlice2)
	copy(newSlice2, slice)
	// [Apple Orange Plum Banana Grape]
	fmt.Println(newSlice2)

	newSlice3 := []string{"Apple2", "Orange3", "Apple"}
	copy(slice, newSlice3)
	fmt.Println(newSlice3)
	fmt.Println(slice)
}
