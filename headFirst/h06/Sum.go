package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// OpenFile 打开文件并返回文件句柄
func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

// CloseFile 关闭文件
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

// GetFloats 读取文件中的数据
func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	CloseFile(file)
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)

}
