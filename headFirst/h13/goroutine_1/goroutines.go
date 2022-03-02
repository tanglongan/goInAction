package main

// goroutine：是同步运行的函数。新的goroutine以一个go语句开始，一个普通的函数调用，前面是关键字“go”
// channel：是用于goroutine之间发送值的数据结构。
// 默认情况下，在channel上发送一个值会阻塞当前goroutine的执行，直到接收到该值为止。试图接收一个值也会阻塞当前goroutine，直到被发送到那个channel为止。
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
