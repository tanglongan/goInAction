package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// scanDirectory 递归获取给定目录下的所有文件和子目录，但是其中只是打印了错误信息，这样并不够好。
func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// 打印ReadDir函数调用中的错误信息。但是这样处理错误的方式并不理想！
		fmt.Printf("Returning error from scanDirectory( \"%s\" )\n", path)
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())

		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				// 打印ReadDir函数调用中的错误信息。但是这样处理错误的方式并不理想！
				fmt.Printf("Returning error from scanDirectory( \"%s\" )\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("D:\\")
	if err != nil {
		log.Fatal(err)
	}
}
