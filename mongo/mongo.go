package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client mongo.Client

var Database mongo.Database

func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb://root:chino@localhost:27017").SetServerAPIOptions(serverAPIOptions)
	_Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = _Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	Client = *_Client
	_Database := Client.Database("rearyard_compose")
	Database = *_Database
	fmt.Println("Connected to MongoDB!")
}
