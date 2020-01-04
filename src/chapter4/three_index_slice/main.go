package main

import "fmt"

func main() {
	// 创建字符串切片
	// 其长度和容量都是5 个元素
	slice := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是1 个元素
	newSlice := slice[2:3:3]

	// 此时slice和newSlice还是共享同一个地址空间
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[2], &newSlice[0])

	newSlice = append(newSlice, "Wiki")
	// 此时已经创建了一个新的底层数组
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[2], &newSlice[0])
	fmt.Printf("newSlice:%v", newSlice)
}
