
proto_test: test/state.proto
	mkdir -p test/go
	protoc --go_opt=paths=source_relative --go_out=test/go test/state.proto
