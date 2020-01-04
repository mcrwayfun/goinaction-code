package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是5个元素
	slice := []int{10, 20, 30, 40, 50}

	// 创建一个新的切片
	// 其长度为2个元素，容量为4个元素
	newSlice := slice[1:3]

	// 输出：2, 4
	fmt.Printf("newSlice len:%d, cap:%d\n", len(newSlice), cap(newSlice))
	// 输出：[20 30]
	fmt.Println(newSlice)

	// error : index out of range
	// fmt.Println(newSlice[2])

	newSlice = append(newSlice, 60)
	newSlice = append(newSlice, 70)
	// 输出：[20 30 60 70 80]
	fmt.Println(newSlice)
	// 输出：[10 20 30 60 70]
	fmt.Println(slice)

}
