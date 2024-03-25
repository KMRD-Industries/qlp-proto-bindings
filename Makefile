
proto_test: test.proto
	mkdir -p test/go test/cpp
	protoc --go_opt=paths=source_relative --go_out=test/go test.proto
	protoc --cpp_out=test/cpp test.proto
