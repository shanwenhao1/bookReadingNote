# mock simple example

To run the test,

```bash
go test bookReadingNote/project/myMock
```

To Generate the `Spider` mock
```bash
# under dir 'bookReadingNote\project\mock\myMock\spider'
mockgen -destination mock_spider/mock_spider.go bookReadingNote/project/mock/myMock/spider Spider
```

## 参考
- [使用Golang的官方mock工具--gomock](https://www.jianshu.com/p/598a11bbdafb)