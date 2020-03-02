package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/micro/go-micro/broker"
	pb "github.com/superryanguo/kick/user_service/proto"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const topic = "user.created"

type service struct {
	repo         Repository
	tokenService Authable
	//Publisher    micro.Publisher
	PubSub broker.Broker
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		res.Done = false
		return err
	}
	res.Done = true
	res.Users = append(res.Users, user)
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		res.Done = false
		return err
	}
	res.Done = true
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Auth user:", req.Email, req.Password)
	user, err := srv.repo.GetByEmail(req.Email)
	log.Printf("Auth find user %v\n", user)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	log.Println("Auth pass")

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil

}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hp)

	if err := srv.repo.Create(req); err != nil {
		res.Done = false
		return err
	}
	res.Done = true
	log.Printf("Create the user%v\n", req)
	res.Users = append(res.Users, req)

	if err := srv.publishEvent(req); err != nil {
		return err
	}
	return nil
}

func (srv *service) publishEvent(user *pb.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	if err := srv.PubSub.Publish(topic, msg); err != nil {
		log.Printf("[broker pub] failed: %v", err)
	}

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println("cliams:", claims)

	if claims.User.Id == "" { //TODO: need to enhance?
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
