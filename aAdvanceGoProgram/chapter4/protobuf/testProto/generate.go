//go:generate protoc --proto_path=. --proto_path=../../../../rpcDepend --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. test.proto
package testProto
