package main

import "fmt"

func main() {
	s := "中国欢迎您"
	// 一个rune对应一个中文字符对应的Unicode码点
	rs := []rune(s)
	//UTF-8编码是Unicode码点的一种字符编码形式，是最常用的一种编码格式，也是Go默认的编码格式
	s1 := []byte(s)

	for i, v := range rs {
		var utf8Bytes []byte

		//UTF-8编码中，大多数中文字符都是用三字节表示。[]byte(s)的转型让我们获得了底层存储的副本，从而得到了每个汉字字符的编码字节
		for j := i * 3; j < (i+1)*3; j++ {
			utf8Bytes = append(utf8Bytes, s1[j])
		}
		// 打印字符、Unicode码点 UTF8编码
		fmt.Printf("%s => %X => %X\n", string(v), v, utf8Bytes)
	}

}
