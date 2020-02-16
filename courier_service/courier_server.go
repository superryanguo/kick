package main

import (
	"context"
	"errors"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/courier_service/proto"
)

type Repoer interface {
	Dispatch(req *pb.Request) (*pb.Courier, error)
	ShowAll() []*pb.Courier
}

type Repo struct {
	couriers []*pb.Courier
}

func (repo *Repo) Dispatch(req *pb.Request) (*pb.Courier, error) {
	for _, cr := range repo.couriers {
		if cr.Available && req.Quantity <= cr.Capacity && (req.Quantity*req.Weight) <= cr.MaxWeight {
			//update the data
			cr.Available = false
			cr.Capacity -= req.Quantity
			cr.MaxWeight -= (req.Quantity * req.Weight)
			cr.OrderId = append(cr.OrderId, req.OrderId)
			log.Printf("Assign %s[%s] for %s\n", cr.Name, cr.CourierId, cr.OrderId)
			return cr, nil
		}
	}
	return nil, errors.New("No Available Courier")
}

func (repo *Repo) ShowAll() []*pb.Courier {
	return repo.couriers
}

type service struct {
	repo Repoer
}

func (s *service) Dispatch(ctx context.Context, req *pb.Request, res *pb.Response) error {

	c, err := s.repo.Dispatch(req)
	if err != nil {
		return err
	}

	res.Dispatched = true
	res.Courier = c
	return nil
}

func (s *service) ShowAll(ctx context.Context, req *pb.ShowRequest, res *pb.Result) error {
	res.Couriers = s.repo.ShowAll()
	return nil
}

func main() {

	repo := &Repo{couriers: make([]*pb.Courier, 3)}
	repo.couriers[0] = &pb.Courier{CourierId: "c001", Name: "Kane", MaxWeight: 200, Capacity: 5, Available: true, OrderId: nil}
	repo.couriers[1] = &pb.Courier{CourierId: "c002", Name: "Peter", MaxWeight: 300, Capacity: 2, Available: true, OrderId: nil}
	repo.couriers[2] = &pb.Courier{CourierId: "c003", Name: "Charles", MaxWeight: 500, Capacity: 3, Available: true, OrderId: nil}

	srv := micro.NewService(
		micro.Name("courier"),
	)

	srv.Init()

	pb.RegisterCourierServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to CourierServer: %v", err)
	}
}
