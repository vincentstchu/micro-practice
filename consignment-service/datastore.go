package main
import (
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

//func CreateSession(host string) (*mongo.Session, error) {
//	client, err := CreateClient(host)
//	if err != nil {
//		return nil, err
//	}
//	session, err := client.StartSession()
//	return session, err
//}

func CreateClient(host string) (*mongo.Client, error)  {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	return client, err
}