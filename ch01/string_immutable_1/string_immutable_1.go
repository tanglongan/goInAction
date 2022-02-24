package main

import "fmt"

func main() {
	// 原始字符串
	var s = "hello"
	fmt.Println("original string:", s)

	// 切片化之后试图改变原字符串
	// 对string进行切片化之后，Go编译器会为切片变量重新分配底层存储而不是公用string的底层存储，因此对切片的修改并未对原string的数据产生任何影响
	s1 := []byte(s)
	s1[0] = 't'
	fmt.Println("slice: ", string(s1))
	fmt.Println("after reslice, the original string is: ", s)
}
