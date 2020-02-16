build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  outlet_service/proto/outlet.proto
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  courier_service/proto/courier.proto
	docker build -t outlet_server  -f ./outlet_service/Dockerfile ./outlet_service
	docker build -t courier_server -f ./courier_service/Dockerfile ./courier_service
	docker build -t outlet_client -f ./outlet_client/Dockerfile ./outlet_client

run:
	docker run -p 7000:7000 outlet_server&
	docker run courier_server&
	docker run outlet_client

