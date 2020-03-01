package main

import (
	"context"
	"errors"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	cor "github.com/superryanguo/kick/courier_service/proto"
	pb "github.com/superryanguo/kick/outlet_service/proto"
	userService "github.com/superryanguo/kick/user_service/proto"
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
		micro.WrapHandler(AuthWrapper),
	)

	c := cor.NewCourierServiceClient("courier", srv.Client())
	srv.Init()

	pb.RegisterOutletServiceHandler(srv.Server(), &service{session, c})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		//skip it when DISABLE_AUTH is set
		if os.Getenv("DISABLE_AUTH") == "true" {
			log.Println("skipping the auth by ENV set")
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"] //why upcase?!
		if len(token) == 0 {
			log.Println("The Token is empty, return")
			return errors.New("empty token")
		} else {
			log.Println("AuthWrapper Authenticating with token: ", token)
		}

		authClient := userService.NewUserServiceClient("user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(ctx, &userService.Token{
			Token: token,
		})
		log.Println("Auth resp:", authResp)
		log.Println("Err:", err)
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
