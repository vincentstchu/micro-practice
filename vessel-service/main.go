package main

import (
	"os"
	"context"
	"log"
	"github.com/micro/go-micro"
	pb "shippy/vessel-service/proto/vessel"
)

const (
	DefaultHost = "mongodb://localhost:27017"
)
func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = DefaultHost
	}
	client ,err := CreateClient(host)
	defer client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("creat client error: &v\n", err)
	}
	repo := &VesselRepository{client}
	CreateDummyData(repo)
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()
	pb.RegisterVesselServiceHandler(server.Server(), &handler{client})
	

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CreateDummyData(repo Repository)  {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}