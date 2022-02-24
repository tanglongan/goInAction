package main

// Go语言直接提供通过反引号构造"所见即所得"的多行字符串
import "fmt"

const a = `红豆生南国，
春来发几枝。
愿君多采撷，
此物最相思。`

func main() {
	fmt.Println(a)
}
