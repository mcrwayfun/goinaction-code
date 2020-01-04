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
	// 输出：[20 30 60 70]
	fmt.Println(newSlice)
	// 输出：[10 20 30 60 70]
	fmt.Println(slice)
	// slice[1]的地址应该与newSlice[0]的一样
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[1], &newSlice[0])

	newSlice = append(newSlice, 80)
	// 输出：[20 30 60 70 80]
	fmt.Println(newSlice)
	// 输出：[10 20 30 60 70]
	fmt.Println(slice)
	fmt.Printf("slice len:%d, cap:%d\n", len(slice), cap(slice))
	fmt.Printf("newSlice len:%d, cap:%d\n", len(newSlice), cap(newSlice))
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[4], &newSlice[4])

	// 创建切片时，新切片的长度和容量与旧的设置成一致
	// 新切片在进行append操作时，会创建一个新的底层数组，并将原值copy到新的数组中
	newSlice2 := slice[0:5]
	// 此时还未进行append操作，两者第一个元素存储地址应该相同
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[0], &newSlice2[0])
	// 输出：[10 20 30 60 70]
	fmt.Println(newSlice2)
	newSlice2 = append(newSlice2, 180)
	// 输出：[10 20 30 60 70 180]
	fmt.Println(newSlice2)
	// 输出两个地址不一致
	fmt.Printf("address of slice:%p, newSlice:%p\n", &slice[0], &newSlice2[0])
}
