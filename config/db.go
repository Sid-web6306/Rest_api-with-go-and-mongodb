package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Sid-web6306/restapi/controllers"
)
// CONNECTION TO MONGODB
func Connect(){
	clientOptions :=options.Client().ApplyURI("mongodb://localhost:27017")

	client,err:= mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	err = client.Connect(ctx)

	defer cancel()
	err = client.Ping(context.Background(),readpref.Primary())
	if err!=nil{
		log.Fatal("Couldn't connect to the database: ",err)

	}else{
		log.Println("Connected to database Successfully")
	}
	db:= client.Database("go_mongodb")
	controllers.PersonCollection(db)

	return

}
