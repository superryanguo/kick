package main

import (
	"context"
	"log"

	cor "github.com/superryanguo/kick/courier_service/proto"
	pb "github.com/superryanguo/kick/outlet_service/proto"
	//"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type service struct {
	session       *mgo.Session
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

func (s *service) GetRepo() Repoer {
	return &Repo{s.session.Clone()}
}

func (s *service) CreateOrder(ctx context.Context, req *pb.Order, res *pb.Response) error {
	defer s.GetRepo().Close()

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

	rt, err := s.courierClient.ShowAll(context.Background(), &cor.ShowRequest{})
	if err != nil {
		return err
	} else {
		for _, v := range rt.Couriers {
			log.Printf("courier:%v\n", v)
		}
	}
	orders, err := s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Orders = orders
	log.Printf("The order %s[%s] is Created!\n", req.Id, req.UserId)
	return nil
}

func (s *service) GetOrders(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer s.GetRepo().Close()
	orders := s.GetRepo().GetAll()
	res.Orders = orders
	return nil
}
