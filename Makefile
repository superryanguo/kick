build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  outlet_service/proto/outlet.proto
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/superryanguo/kick \
	  courier_service/proto/courier.proto
	docker build -t outlet_server ./outlet_service/Dockerfile
	docker build -t courier_server ./courier_service/Dockerfile
	docker build -t outlet_client ./outlet_client/Dockerfile

run:
	docker run -p 7000:7000 outlet_server
	docker run -p 7001:7001 courier_server
	docker run outlet_client

