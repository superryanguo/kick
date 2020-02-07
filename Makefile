build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  outlet_service/proto/outlet.proto
