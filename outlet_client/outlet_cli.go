package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/superryanguo/kick/outlet_service/proto"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:7000"
	defaultFilename = "order.json"
)

func parseFile(file string) (*pb.Order, error) {
	var order *pb.Order
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &order)
	return order, err
}

func main() {
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//log.Fatalf("Did not connect: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewOutletServiceClient(conn)

	service := micro.NewService(micro.Name("outlet"))
	service.Init()

	client := pb.NewOutletServiceClient("outlet", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	order, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetOrders(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list orders: %v", err)
	}
	for _, v := range getAll.Orders {
		log.Println(v)
	}
}
