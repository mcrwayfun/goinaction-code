package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt通道报告从操作系统发送信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	timeout <-chan time.Time

	// tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeOut会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将一个任务附加到Runner上
// 这个任务是一个接收int类型的ID
// 作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 希望接收所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)
	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已经注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任务信号
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}

// timeout 规定了必须在多少秒内完成
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work")

	// 为本次执行分配超时时间
	r := New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}

// createTask 返回一个根据id
// 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
