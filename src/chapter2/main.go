package main

import (
	_ "chapter2/matchers"
	"chapter2/search"
	"log"
	"os"
)

// init在main之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main 是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("UNICEF")
}
