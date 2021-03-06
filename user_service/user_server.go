package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/user_service/proto"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		log.Println("connecting to the database is ok")
	}

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)

	srv.Init()

	//publisher := micro.NewPublisher("user.created", srv.Client())
	// Get instance of the broker using our defaults
	pubsub := srv.Server().Options().Broker

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, pubsub})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
