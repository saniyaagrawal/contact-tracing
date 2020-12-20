package main

import (
	"fmt"
	"net/http"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty`
	Name string `json:"_id, omitempty" bson:"_id, omitempty`
	// Date of Birth
	// Phone Number
	// Email Address
	// Creation Timestamp

}

func newUserHandler(response http.ResponseWriter, request *http.Request){
	fmt.Print("post api for users")
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	
	err = client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())

	http.HandleFunc("/users",newUserHandler)

	err = http.ListenAndServe(":8080", nil)
	if(err!=nil){
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}