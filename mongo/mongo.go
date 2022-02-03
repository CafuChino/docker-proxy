package mongo

import (
	"docker-controller/conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client mongo.Client

var Database mongo.Database

func init() {
	conf.LoadConfig("");
	config := conf.Conf.Mongo;
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)).SetServerAPIOptions(serverAPIOptions)
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
