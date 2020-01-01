package search

import "log"

// Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的
// 新定义类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine 来执行这个函数
// 并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Panicln(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}
