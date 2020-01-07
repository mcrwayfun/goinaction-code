package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// admin代表一个拥有权限的管理员用户
type admin struct {
	user  // 嵌入类型
	level string
}

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// notify 使用指针接收者实现了notifier 接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// 创建一个admin用户
	ad := admin{
		user: user{
			name:  "mcrwayfun",
			email: "mcrwayfun@gmail.com",
		},
		level: "super",
	}

	// 可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法也被是提升到外部类型
	ad.notify()
}
