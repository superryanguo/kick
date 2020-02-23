package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/courier_service/proto"
)

const (
	defaultDb = "localhost:27017"
)

func createDummyData(repo Repoer) {
	defer repo.Close()
	couriers := make([]*pb.Courier, 3)
	couriers[0] = &pb.Courier{CourierId: "c001", Name: "Kane", MaxWeight: 200, Capacity: 5, Available: true, OrderId: nil}
	couriers[1] = &pb.Courier{CourierId: "c002", Name: "Peter", MaxWeight: 300, Capacity: 2, Available: true, OrderId: nil}
	couriers[2] = &pb.Courier{CourierId: "c003", Name: "Charles", MaxWeight: 500, Capacity: 3, Available: true, OrderId: nil}
	for _, v := range couriers {
		repo.Create(v)
	}
}
func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultDb
	}

	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}
	repo := &Repo{session.Copy()}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("courier"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterCourierServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to CourierServer: %v", err)
	}
}
