# Samples

This directory contains an example of a package containing a non-trivial
interface that can be mocked with GoMock. The interesting files are:

* `user.go`: Source code for the sample package, containing interfaces to be
    mocked. This file depends on the packages named imp[1-4] for various things.

* `user_test.go`: A test for the sample package, in which mocks of the
    interfaces from `user.go` are used. This demonstrates how to create mock
    objects, set up expectations, and so on.

* `mock_user/mock_user.go`: The generated mock code. See ../gomock/matchers.go
    for the `go:generate` command used to generate it.

To run the test,

```bash
go test bookReadingNote/project/mockSample
```

To Generate the mock code
```bash
mockgen -destination mock_user/mock_user.go bookReadingNote/microservicesPatterns/code/chapter3/mockSample Index,Embed,Embedded
```

## 更简单的mock[示例](../myMock/README.md)
