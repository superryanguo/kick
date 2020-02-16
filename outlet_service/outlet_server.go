package main

import (
	"log"

	micro "github.com/micro/go-micro"
	cor "github.com/superryanguo/kick/courier_service/proto"
	pb "github.com/superryanguo/kick/outlet_service/proto"
	"golang.org/x/net/context"
)

const (
	port = ":7000"
)

type Repoer interface {
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
	repo          Repoer
	courierClient cor.CourierServiceClient
}

func GetQuantity(c []*pb.Commodity) int32 {
	var q int32
	for _, d := range c {
		q += d.Quantity
	}
	return q
}

func GetWeight(c []*pb.Commodity) int32 {
	var w int32
	for _, d := range c {
		w += (d.Weight * d.Quantity)
	}
	return w
}
func (s *service) CreateOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {

	cr, err := s.courierClient.Dispatch(context.Background(), &cor.Request{
		Quantity: GetQuantity(req.Commoditys),
		Weight:   GetWeight(req.Commoditys),
		OrderId:  req.Id,
	})
	if err != nil {
		return err
	}
	//update the data into order
	if cr.Dispatched {
		req.CourierId = cr.Courier.CourierId
		log.Printf("Found courier: %s[%s]\n", cr.Courier.Name, req.CourierId)
	} else {
		log.Println("The courier is not Available!")
	}

	orders, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Orders = orders
	log.Printf("The order %s[%s] is Created!\n", req.Id, req.UserId)
	return nil
}

func (s *service) GetOrders(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	orders := s.repo.GetAll()
	res.Orders = orders
	return nil
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

	c := cor.NewCourierServiceClient("courier", srv.Client())
	srv.Init()

	pb.RegisterOutletServiceHandler(srv.Server(), &service{repo, c})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
