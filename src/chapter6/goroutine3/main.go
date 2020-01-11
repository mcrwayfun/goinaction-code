package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter增加包里counter变量的值
func incCounter(id int) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获counter的值
		value := counter

		// 当前goroutine从线程退出时，并返回队列
		runtime.Gosched()

		// 增加本地value变量的值
		value++

		// 将该值保存回counter
		counter = value
	}
}
