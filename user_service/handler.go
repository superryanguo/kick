package main

import (
	pb "github.com/superryanguo/kick/user_service/proto"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		res.Done = false
		return err
	}
	res.Done = true
	res.Users[0] = user
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
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		res.Done = false
		return err
	}
	res.Done = true
	res.Users[0] = req
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
