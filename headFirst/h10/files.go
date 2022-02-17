package main

// 使用了panic，简化了错误处理代码。但是panic同样会导致程序崩溃，出现难看的堆栈跟踪。
import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// 当读取目录遇到错误时，就产生一个panic，程序退出并且打印错误栈信息
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

// 报告panic
func reportPanic() {
	// 调用recover()函数并存储它的返回值
	p := recover()
	if p == nil {
		return
	}
	// 从recover返回值中获取底层数据
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	defer reportPanic()
	scanDirectory("D:\\")
}
