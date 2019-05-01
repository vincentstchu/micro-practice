package main

import (
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	ADDRESS = "localhost:50051"
	DEFAULT = "consignment.json"
)

func parseFile(path string) (*pb.Consignment, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(raw, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file unmarshal error")
	}
	return consignment, nil
}

func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)

	file := DEFAULT
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
