package myMock

import (
	"bookReadingNote/project/mock/myMock/spider/mock_spider"
	"github.com/golang/mock/gomock"
	"testing"
)

// mock test GetGoVersion
func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := mock_spider.NewMockSpider(mockCtl)
	// mock GetBody return
	mockSpider.EXPECT().GetBody().Return("go1.8.3")

	// begin test
	goVer := GetGoVersion(mockSpider)
	if goVer != "go1.8.3" {
		t.Error("Get wrong version: ", goVer)
	}
}
