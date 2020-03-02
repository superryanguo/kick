package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/user_service/proto"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("->Got a message..........")
	log.Println("Sending email to...:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("emailsrv"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
