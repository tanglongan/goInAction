package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer值，并通过io.Writer的Write方法将一个字符串写入Buffer
	var b bytes.Buffer
	b.Write([]byte("hello"))

	// 使用Fprintf将一个字符串拼接到Buffer里
	_, err := fmt.Fprintf(&b, " World!")
	if err != nil {
		return
	}

	// 将Buffer的内容输出到标准输出设备
	_, err = b.WriteTo(os.Stdout)
	if err != nil {
		return
	}
}
