package main

import "fmt"

func main() {
	rs := []rune{0x4E2D,
		0x56FD,
		0x6B22,
		0x8FCE,
		0x60A8,
	}
	// []rune转化为string
	s := string(rs)
	fmt.Println(s)

	s1 := []byte{
		0xE4, 0xB8, 0xAD,
		0xE5, 0x9B, 0xBD,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
	}
	// []byte转化为string
	s2 := string(s1)
	fmt.Println(s2)

	ch := "中国欢迎您"
	// 一个rune对应一个中文字符对应的Unicode码点
	ru := []rune(ch)
	fmt.Println(ru)
	//UTF-8编码是Unicode码点的一种字符编码形式，是最常用的一种编码格式，也是Go默认的编码格式
	by := []byte(ch)
	fmt.Println(by)

	for i, v := range ru {
		var utf8Bytes []byte

		//UTF-8编码中，大多数中文字符都是用三字节表示。[]byte(s)的转型让我们获得了底层存储的副本，从而得到了每个汉字字符的编码字节
		for j := i * 3; j < (i+1)*3; j++ {
			utf8Bytes = append(utf8Bytes, by[j])
		}
		// 打印字符、Unicode码点 UTF8编码
		fmt.Printf("%s => %X => %X\n", string(v), v, utf8Bytes)
	}
}
