package main

import (
	"context"

	pb "github.com/superryanguo/kick/courier_service/proto"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
	//TODO: change to the repoer?!
}

func (s *service) GetRepo() Repoer {
	return &Repo{s.session.Clone()}
}

func (s *service) Dispatch(ctx context.Context, req *pb.Request, res *pb.Response) error {

	defer s.GetRepo().Close()
	c, err := s.GetRepo().Dispatch(req)
	if err != nil {
		return err
	}

	res.Dispatched = true
	res.Courier = c
	return nil
}

func (s *service) ShowAll(ctx context.Context, req *pb.ShowRequest, res *pb.Result) error {
	defer s.GetRepo().Close()
	res.Couriers = s.GetRepo().ShowAll()
	return nil
}

func (s *service) Create(ctx context.Context, c *pb.Courier, res *pb.Response) error {
	defer s.GetRepo().Close()
	res.Dispatched = false
	res.Courier = c
	return s.GetRepo().Create(c)

}
