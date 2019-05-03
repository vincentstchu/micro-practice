package main

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	pb "shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName        = "shippy"
	ConCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	client *mongo.Client
}

func (cr *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	_, err := cr.getMongoCollection().InsertOne(context.TODO(), consignment)
	return err
}

func (cr *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var cons []*pb.Consignment
	findOptions := options.Find()
	findOptions.SetLimit(3)
	cursor, err := cr.getMongoCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var item *pb.Consignment
		err = cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		cons = append(cons, item)
	}
	cursor.Close(context.TODO())
	return cons, err
}

//TODO: not functional right now
func (cr *ConsignmentRepository) Close() {
	// cr.client.Disconnect(context.TODO())
}

func (cr *ConsignmentRepository) getMongoCollection() *mongo.Collection{
	return cr.client.Database(DbName).Collection(ConCollection)
}