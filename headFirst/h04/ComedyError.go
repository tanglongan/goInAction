package main

import "fmt"

type ComeError string

// Error 满足error接口
func (c ComeError) Error() string {
	// Error接口需要返回string，因此需要做类型转换
	return string(c)
}

func main() {
	var err error
	err = ComeError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)
}
