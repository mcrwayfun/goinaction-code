package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d Index-Addr: %X Value-Addr: %X ElemAddr: %X\n",
			index, value, &index, &value, &slice[index])
	}

	for index, value := range slice {
		slice[index] = value + 10
		fmt.Printf("Index: %d Value: %d Index-Addr: %X Value-Addr: %X ElemAddr: %X ElemValue:%d\n",
			index, value, &index, &value, &slice[index], slice[index])
	}
	fmt.Println(slice)

	for index := 0; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d Index-Addr: %X Value-Addr: %X\n",
			index, slice[index], &index, &slice[index])
	}
}
