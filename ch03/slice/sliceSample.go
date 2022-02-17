package slice

import "fmt"

func test() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	// append函数会向切片中追加元素，总是会增加新切片的长度，而容量可能增加也可能不会改变，这取决于被操作切片的可用容量
	newSlice = append(newSlice, 60)
	fmt.Println(newSlice)
}
