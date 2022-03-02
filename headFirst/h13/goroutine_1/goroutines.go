package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	resultChannel := make(chan Page)
	urls := []string{"https://www.baidu.com", "https://www.google.com.hk", "https://www.sogou.com"}

	for _, url := range urls {
		go responseSize(url, resultChannel)
	}

	for i := 0; i < len(urls); i++ {
		page := <-resultChannel
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	fmt.Println("GET ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 通过channel发送网页结果对象
	channel <- Page{URL: url, Size: len(body)}
}
