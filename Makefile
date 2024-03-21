
proto_test: test.proto
	mkdir -p test/go
	protoc --go_opt=paths=source_relative --go_out=test/go test.proto
