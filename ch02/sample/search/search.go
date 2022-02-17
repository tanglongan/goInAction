package search

import (
	"log"
	"sync"
)

// 注册用户搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchTerm string) {
	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	//创建一个无缓冲的通道，接收匹配后的结果（通道本身实现的是一组带类型的值，这组值用于在goroutine之间传递数据，通道内置同步机制，从而保证通信安全）
	results := make(chan *Result)

	//创建一个waitGroup，跟踪所有启动的goroutine是否完成工作（WaitGroup是一个计数信号量，用它可以统计所有goroutine是不是已经都完成工作了）
	var waitGroup sync.WaitGroup

	//设置需要等待处理
	//将waitGroup中信号量设置为将要启动的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找特定数据源类型
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个goroutine来执行搜索，每个数据源对应一个goroutine
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		// 等待所有任务完成（waitGroup的wait()方法会导致goroutine阻塞，直到WaitGroup内部的计数到达0时，当前goroutine继续向下执行）
		waitGroup.Wait()
		// 调用内置close()函数，关闭通道
		close(results)
	}()

	// 启动函数，显示返回的结果
	// 在最后一个结果显示完成后返回，返回后当前程序就会终止
	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	// 将一个Matcher值加入到保存注册匹配器的映射中
	matchers[feedType] = matcher
}
