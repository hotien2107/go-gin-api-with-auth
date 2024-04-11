package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB *mongo.Database
}

var mongoDatabase *mongo.Database

func NewMongoDB() *MongoDB {
	if mongoDatabase == nil {
		mongoDatabase = &mongo.Database{}
	}
	return &MongoDB{DB: mongoDatabase}
}

func (mongoDB *MongoDB) InitMongodb() {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	mongoDatabase = client.Database("chat-app")
	fmt.Println("Connect to mongodb successfully!")
}
