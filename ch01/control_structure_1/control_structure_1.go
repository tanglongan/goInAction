package main

// range表达式的复制行为会带来性能上的消耗：
// 1、当range表达式的类型是数组时，range需要复制整个数组
// 2、当range表达式的类型时切片时，这个消耗就小很多，仅仅是复制一个指针或切片的内部表示（一个结构体）
// 3、range表达式对于string、map、channel都会生成一个副本数据
import (
	"fmt"
)

func sliceRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("sliceRangeExpression result:")
	fmt.Println("a= ", a)

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r= ", r)
	fmt.Println("a= ", a)
}

// sliceLenChangeRangeExpression 切片和数据有一个不同点：就是切片的len在运行时可以被改变，而数组长度是一个常量，不允许改变
func sliceLenChangeRangeExpression() {
	a := []int{1, 2, 3, 4, 5}
	r := make([]int, 0)

	fmt.Println("sliceLenChangeRangeExpression result: ")
	fmt.Println("a= ", a)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)

	}

	fmt.Println("r= ", r)
	fmt.Println("a= ", a)
}

func main() {
	sliceRangeExpression()
	fmt.Println("======================")
	sliceLenChangeRangeExpression()
}
