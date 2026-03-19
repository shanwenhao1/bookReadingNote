//go:generate protoc --proto_path=./ --go_out=paths=source_relative:./ esconf.proto
// protoc手动生成
//protoc --proto_path=. --proto_path=../../../../third_party --proto_path=../../../ --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --validate_out=paths=source_relative,lang=go:. ibintegral.proto

package esconf
