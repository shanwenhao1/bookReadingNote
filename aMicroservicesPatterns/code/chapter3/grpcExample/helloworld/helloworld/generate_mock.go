//go:generate mockgen -destination ../mock_helloworld/my_mock.go bookReadingNote/microservicesPatterns/code/chapter3/grpcExample/helloworld/helloworld/google.golang.org/grpc/examples/helloworld/helloworld GreeterClient
package helloworld
