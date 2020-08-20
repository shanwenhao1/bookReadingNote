package myMock

import "bookReadingNote/project/mock/myMock/spider"

// 测试GetGoVersion函数, 使用mock实现跳过Spider依赖进行测试(因此不需要实现Spider 接口的GetBody函数从而
// 测试GetGoVersion函数)
func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
