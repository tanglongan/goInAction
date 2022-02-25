package main

import "fmt"

// 按照普通求值规则，这些函数调用，channel操作从左向右的顺序进行求值
// 当普通求值顺序与与包级变量依赖顺序一并使用时，后者优先级更高，但单个单独表达式中操作数求值依旧按照普通求值顺序的规则

var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int {
	fmt.Println("calling f()")
	return c
}

func g() int {
	fmt.Println("calling g()")
	return 1
}

func sqr(x int) int {
	fmt.Println("calling sqr()")
	return x * x
}

func v() int {
	fmt.Println("calling v()")
	return 1
}

func u() int {
	fmt.Println("calling u()")
	return 2
}

func main() {
	// 6、1、5
	fmt.Println(a, b, c)
}
