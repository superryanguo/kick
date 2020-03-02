package main

import (
	"context"
	"encoding/json"
	"log"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
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

	// Get the broker instance using our environment variables
	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	//micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	// Subscribe to messages on the broker
	_, err := pubsub.Subscribe(topic, func(p broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
func sendEmail(user *pb.User) error {
	log.Println("->Got a message..........")
	log.Println("Sending email to...:", user.Name)
	return nil
}
