//go:generate mockgen -destination mock_spider/mock_spider.go bookReadingNote/project/mock/myMock/spider Spider
package spider

type Spider interface {
	GetBody() string
}
