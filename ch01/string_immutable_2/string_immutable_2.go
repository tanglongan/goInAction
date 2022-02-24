package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 原始string
	s := "hello"
	fmt.Println("Original string:", s)

	// 试图通过unsafe指针改变原始string
	modifyString(&s)
	fmt.Println(s)
}

// 通过unsafe指针指向string在运行时内部表示结构中的存储块的地址，然后通过指针修改那块内存中存储的数据。
// 程序运行可以看到，对string的底层的数据存储仅能进行只能读操作，一旦试图修改那块数据，就会得到SIGBUS运行时错误。
func modifyString(s *string) {
	// 取出第一个8字节的值
	p := (*uintptr)(unsafe.Pointer(s))
	// 获取底层数组的地址
	array := (*[5]byte)(unsafe.Pointer(*p))

	len := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		// 尝试修改字符
		(*p1) = v + 1
	}
}
