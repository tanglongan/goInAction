package main

import "fmt"

// notifier 是一个通知类行为的接口
type notifier interface {
	notify()
}

// user 用户类型
type user struct {
	name  string
	email string
}

// notify 实现了一个可以通过user类型值指针
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin 拥有权限的管理员用户类型
type admin struct {
	user  // 嵌入类型
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	ad := admin{
		user:  user{name: "smith", email: "smith@gmail.com"},
		level: "super",
	}

	// admin 实现了notifier接口，因此覆盖了内部类型user相关的同名方法
	sendNotification(&ad)
	// 通过内部属性访问内部属性的方法
	ad.user.notify()

	// 内部类型的方法也被提升到外部类型
	ad.notify()
}

// 发送通知
func sendNotification(n notifier) {
	n.notify()
}
