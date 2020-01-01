package search

// defaultMatcher 实现了默认的匹配器
type defaultMatcher struct{}

// init 函数将默认的匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
func (d defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
