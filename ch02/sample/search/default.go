package search

// defaultMatcher 实现了默认匹配器（空结构体在创建实力时，不会分配任何内存，很适合创建没有任何状态的类型）
type defaultMatcher struct {
}

// init函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
