# protobuf code-gen command
gen:
	protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. ./rpc/service.proto

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

