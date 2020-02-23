package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	cor "github.com/superryanguo/kick/courier_service/proto"
	pb "github.com/superryanguo/kick/outlet_service/proto"
)

const (
	defaultDb = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultDb
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	//repo := &Repo{}

	// Set-up our gRPC server.
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()

	srv := micro.NewService(
		micro.Name("outlet"),
		micro.Version("latest"),
	)

	c := cor.NewCourierServiceClient("courier", srv.Client())
	srv.Init()

	pb.RegisterOutletServiceHandler(srv.Server(), &service{session, c})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
