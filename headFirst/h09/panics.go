package main

// 当程序出现panic时，所有延迟的函数调用仍然会被执行。
// 如果有多个延迟调用，他们的延迟调用与被延迟的顺序相反。
import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("defered in one()")
	two()
}

func two() {
	defer fmt.Println("defered in two()")
	three()
}

func three() {
	defer fmt.Println("defered in three()")
	panic("This call stack's too deep for me")
}
