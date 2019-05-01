package main

// import (
// 	"context"
// 	"github.com/pkg/errors"
// 	pb "shippy/vessel-service/proto/vessel"
// 	"github.com/micro/go-micro"
// 	"github.com/micro/go-log"
// )

// type Repository interface {
// 	FindAvailable(*pb.Specification) (*pb.Vessel, error)
// }

// type VesselRepository struct {
// 	vessels []*pb.Vessel
// }

// func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
// 	for _, v:= range repo.vessels {
// 		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
// 			return v, nil
// 		}
// 	}
// 	return nil, errors.New("No vessel can't be use")
// }

// type service struct {
// 	repo Repository
// }

// func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
// 	v, err := s.repo.FindAvailable(spec)
// 	if err != nil {
// 		return nil
// 	}
// 	resp.Vessel  = v
// 	return nil
// }

// func main() {
// 	vessels := []*pb.Vessel {
// 		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
// 	}
// 	repo := &VesselRepository{vessels}
// 	server := micro.NewService(
// 		micro.Name("go.micro.serv.vessel"),
// 		micro.Version("latest"),
// 	)
// 	server.Init()
// 	pb.RegisterVesselServiceHandler(server.Server(), &service{repo})
// 	if err := server.Run(); err != nil {
// 		log.Fatal("failed to serve: %v", err)
// 	}
// }

func main() {

}
