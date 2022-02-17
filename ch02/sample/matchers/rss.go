package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"net/http"
	"regexp"
)

// 为了让程序使用文档里的数据，解码RSS文档的时候，需要用到4个数据类型
type (
	// item根据item字段的标签，将定义的字段与rss文档的字段进行关联起来
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image结构与文档的image标签映射关联
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel结构与文档的channel标签字段映射关联
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument定义了与rss文档关联的字段
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// RSS 匹配器结构
type rssMatcher struct {
}

// init函数将匹配器注册到程序里
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search 在文档查找特定的搜索项
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {

	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// 获取要搜索的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 检查标题部分是否包含搜索项，如果包含就添加到结果切片中
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			// append内置函数会根据切片需要，自动对切片进行长度和容量的扩容
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 检查描述部分是否包含搜索项，如果包含就添加到结果切片中
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, err
}

// retrieve 发送HTTP请求获取RSS数据源并解码（该方法只能在包内调用，并未导出）
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// 从网络获取rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	// 一旦当前函数返回，就关闭HTTP响应连接
	defer resp.Body.Close()

	// 检查HTTP响应状态码是不是200，这样就能直到是不是收到了正确的响应
	if resp.StatusCode != 200 {
		// 如果响应码不是200，使用fmt包的Errorf返回一个自定义的错误
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}
	// 将rss数据源文档解码到我们定义的结构类型里，不需要检查错误，调用者会做这件事
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
