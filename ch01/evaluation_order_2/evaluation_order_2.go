package main

// 按照普通求值规则，这些函数调用，channel操作从左向右的顺序进行求值
// 当普通求值顺序与与包级变量依赖顺序一并使用时，后者优先级更高，但单个单独表达式中操作数求值依旧按照普通求值顺序的规则
import (
	"fmt"
)

func f() int {
	fmt.Println("calling f()")
	return 1
}

func g(a, b, c int) int {
	fmt.Println("calling g()")
	return 2
}

func h() int {
	fmt.Println("calling h()")
	return 3
}

func i() int {
	fmt.Println("calling i()")
	return 1
}

func j() int {
	fmt.Println("calling j()")
	return 1
}

func k() bool {
	fmt.Println("calling k()")
	return true
}

func main() {
	var y = []int{11, 12, 13}
	var x = []int{21, 22, 23}
	var c chan int = make(chan int)
	go func() {
		c <- 1
	}()

	// 从左向右的规则，这里的求值顺序是：f()-->h()-->i()-->j()-->c取值操作-->g()-->k()
	y[f()], _ = g(h(), i()+x[j()], <-c), k()
}
