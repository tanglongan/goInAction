package main

import (
	// 让Go语言对包做初始化操作，但是并不使用包里的标识符 。Go编译器不允许声明导入某个包却不使用，下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的init函数
	// 对于这个程序来说，这样做的目的是调用matchers包中的rss.go代码文件里的init函数，注册rss匹配器，以便后用
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init方法在main执行之前调用
func init() {
	// 将日志输出从默认的标准错误stderr，设置为标准输出stdout设备
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	search.Run("president")
}
