package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	pb "shippy/vessel-service/proto/vessel"
)

type handler struct {
	client *mongo.Client
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.client}
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	defer h.GetRepo().Close()
	v, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}