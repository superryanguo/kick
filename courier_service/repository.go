package main

import (
	"errors"
	"log"

	pb "github.com/superryanguo/kick/courier_service/proto"
	"gopkg.in/mgo.v2"
)

const (
	dbName            = "kick"
	courierCollection = "couriers"
)

type Repoer interface {
	Dispatch(req *pb.Request) (*pb.Courier, error)
	ShowAll() []*pb.Courier
	Create(c *pb.Courier) error
	Close()
}

type Repo struct {
	session *mgo.Session
}

func (repo *Repo) Dispatch(req *pb.Request) (*pb.Courier, error) {
	for _, cr := range repo.ShowAll() {
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
	//TODO: need to update to database, not just mem
	return nil, errors.New("No Available Courier")
}

func (repo *Repo) ShowAll() []*pb.Courier {
	var couriers []*pb.Courier
	err := repo.collection().Find(nil).All(&couriers)
	if err != nil {
		log.Fatalf("can't get the data of couriers %s\n", err)
	}
	return couriers
}
func (repo *Repo) Create(c *pb.Courier) error {
	return repo.collection().Insert(c)
}

func (repo *Repo) Close() {
	repo.session.Close()
}

func (repo *Repo) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(courierCollection)
}
