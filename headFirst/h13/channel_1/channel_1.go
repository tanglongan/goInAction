package main

import (
	"fmt"
	"time"
)

func send(msg string, myChannel chan string) {
	fmt.Println("准备发送: ", msg)
	myChannel <- msg
}

func recv(myChannel chan string) {
	msg := <-myChannel
	fmt.Println("接收到值: ", msg)
}

func main() {
	myChannel := make(chan string)
	//在新的goroutine中通过channel发送数据
	go send("hello world", myChannel)
	go recv(myChannel)
	time.Sleep(5 * time.Second)
}
