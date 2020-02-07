package main

import (
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/outlet_service/proto"
	"golang.org/x/net/context"
)

const (
	port = ":7000"
)

type Repoiter interface {
	Create(*pb.Order) ([]*pb.Order, error)
	GetAll() []*pb.Order
}

type Repo struct {
	orders []*pb.Order
}

func (repo *Repo) Create(order *pb.Order) ([]*pb.Order, error) {
	updated := append(repo.orders, order)
	repo.orders = updated
	return updated, nil
}

func (repo *Repo) GetAll() []*pb.Order {
	return repo.orders
}

type service struct {
	repo Repoiter
}

func (s *service) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Response, error) {

	orders, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Orders: orders}, nil
}

func (s *service) GetOrders(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	orders := s.repo.GetAll()
	return &pb.Response{Orders: orders}, nil
}

func main() {

	repo := &Repo{}

	// Set-up our gRPC server.
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()

	srv := micro.NewService(
		micro.Name("outlet"),
	)

	srv.Init()

	pb.RegisterOutletServiceServer(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
