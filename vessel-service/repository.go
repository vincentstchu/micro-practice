package main

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	pb "shippy/vessel-service/proto/vessel"
)

const (
	DbName           = "vessels"
	VesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	client *mongo.Client
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var v *pb.Vessel
	filter := bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}
	err := repo.getMongoCollection().FindOne(context.TODO(), filter).Decode(&v)
	if err != nil {
		log.Println("[x]vessel-service[repo] FindOne ErrorInfo: ", err)
		return nil, err
	}
	return v, err
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	_, err := repo.getMongoCollection().InsertOne(context.TODO(), v)
	return err
}

func (repo *VesselRepository) Close() {
	// repo.client.Disconnect(context.TODO())
}

func (repo *VesselRepository) getMongoCollection() *mongo.Collection {
	return repo.client.Database(DbName).Collection(VesselCollection)
}
