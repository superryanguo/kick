build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/superryanguo/kick \
	  outlet_service/proto/outlet.proto
