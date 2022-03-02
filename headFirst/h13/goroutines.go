package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://www.baidu.com/")
	go responseSize("https://www.google.com.hk/")
	go responseSize("https://www.sogou.com/")

	time.Sleep(5 * time.Second)
	fmt.Println("main goroutine")
}

func responseSize(url string) {
	fmt.Println("getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}
