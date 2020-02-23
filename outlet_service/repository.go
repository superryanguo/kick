package main

import (
	"log"

	pb "github.com/superryanguo/kick/outlet_service/proto"
	"gopkg.in/mgo.v2"
)

const (
	dbName          = "kick"
	orderCollection = "outlet_orders"
)

type Repoer interface {
	Create(*pb.Order) ([]*pb.Order, error)
	GetAll() []*pb.Order
	Close()
}

type Repo struct {
	session *mgo.Session
}

func (repo *Repo) Create(order *pb.Order) ([]*pb.Order, error) {
	err := repo.collection().Insert(order)
	//TODO: error handling
	return repo.GetAll(), err
}

func (repo *Repo) GetAll() []*pb.Order {
	var orders []*pb.Order
	err := repo.collection().Find(nil).All(&orders) // use a simple solution
	if err != nil {
		log.Fatalf("can't get the data of orders %s\n", err)
	}
	return orders

}

func (repo *Repo) Close() {
	repo.session.Close()
}

func (repo *Repo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(orderCollection)
}
