package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
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

	service := micro.NewService(micro.Name("outlet-cli"))
	service.Init()

	client := pb.NewOutletServiceClient("outlet", service.Client())

	file := defaultFilename
	var token string
	if len(os.Args) > 1 {
		file = os.Args[1]
		token = os.Args[2]
	}

	order, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	log.Printf("######let's call CreateOrder\n")
	_, err = client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("Could not CreateOrder: %v", err)
	}

	log.Printf("######let's call GetOrders\n")
	getAll, err := client.GetOrders(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list orders: %v", err)
	}
	for _, v := range getAll.Orders {
		log.Println(v)
	}
}
