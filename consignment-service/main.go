package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
	userPb "shippy/user-service/proto/user"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const (
	DefaultHost = "mongodb://127.0.0.1:27017"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DefaultHost
	}
	client, err := CreateClient(dbHost)
	defer client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		// for auth
		micro.WrapHandler(AuthWrapper),
	)

	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{client, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
