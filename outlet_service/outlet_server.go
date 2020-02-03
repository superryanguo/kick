package main

import (
	"log"
	"net"

	pb "github.com/superryanguo/kick/outlet_service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":7000"
)

type Repoiter interface {
	Create(*pb.Order) (*pb.Order, error)
	GetAll() []*pb.Order
}

type Repo struct {
	orders []*pb.Order
}

func (repo *Repo) Create(order *pb.Order) (*pb.Order, error) {
	updated := append(Repo.orders, order)
	repo.orders = updated
	return orders, nil
}

func (repo *Repo) GetAll() []*pb.Order {
	return Repo.orders
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
