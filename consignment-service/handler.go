package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	// "github.com/micro/go-micro"
	"log"
	go_micro_srv_vessel "shippy/vessel-service/proto/vessel"
	pb "shippy/consignment-service/proto/consignment"
)

type handler struct {
	client *mongo.Client
	vessel go_micro_srv_vessel.VesselServiceClient
}

func (h *handler)GetRepo()Repository  {
	return &ConsignmentRepository{h.client}
}

func (h *handler)CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	defer h.GetRepo().Close()

	log.Println("start to create consignment")
	// vesselService := micro.NewService(micro.Name("micro.srv.vessel"))
	// vesselService.Init()
	// vessel := go_micro_srv_vessel.NewVesselService("micro.srv.vessel", vesselService.Client())
	vReq := &go_micro_srv_vessel.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := h.vessel.FindAvailable(context.Background(), vReq)
	if err != nil {
		log.Println("[x]cosignment-service[handler] vessel.FindAvailable ErrorInfo: ", err)
		return err
	}
	log.Printf("found vessel: %sn", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	err = h.GetRepo().Create(req)
	if err != nil {
		log.Println("[x]cosignment-service[handler] Consignment Create ErrorInfo: ", err)
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (h *handler)GetConsignment(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}