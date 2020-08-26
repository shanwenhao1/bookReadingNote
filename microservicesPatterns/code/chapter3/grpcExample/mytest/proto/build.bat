// generate rpc proto file of test
protoc --go_out=plugins=grpc:. test.proto
protoc --go_out=plugins=grpc:. ./test.proto

// generate mock file of HelloServiceClient interface
mockgen -destination mock_mytest/my_mock.go bookReadingNote/microservicesPatterns/code/chapter3/grpcExample/mytest/proto/test HelloServiceClient