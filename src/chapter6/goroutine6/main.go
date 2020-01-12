package main

import (
	"fmt"
	"sync"
	"time"
)

// 这个程序展示如何用无缓冲的通道来模拟
// 4个goroutine间的接力比赛

var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一位跑步者计数加1
	wg.Add(1)
	// 第一位跑步者持有接力棒
	go Runner(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wg.Wait()
}

// Runner模拟接力赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒
	runner := <-baton
	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton) // 此时下一位跑步者并不会马上开始
	}

	// 围绕跑道跑
	time.Sleep(10 * time.Millisecond)

	// 比赛结束了吗
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchanage With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
