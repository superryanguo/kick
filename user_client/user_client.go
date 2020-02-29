package main

import (
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/config/cmd"
	pb "github.com/superryanguo/kick/user_service/proto"
	"golang.org/x/net/context"
)

func main() {

	//cmd.Init()//TODO: interesting problem with nil para

	client := pb.NewUserServiceClient("user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Name("user-client"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	service.Init(

		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")
			log.Printf("the parameter=%s,%s,%s,%s\n", name, email, password, company)
			if len(name) == 0 || len(password) == 0 {
				log.Println("can't use empty name and mail to create user")
				os.Exit(0)
			}

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %v\n", r.Users[0].Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			// let's just exit because
			os.Exit(0)
		}),
	)

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
