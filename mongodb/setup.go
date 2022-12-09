package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("DB Ping response ok")
	return nil
}

var Client *mongo.Client

func Setup() {
	count := 0
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
label:
	ConErr := Ping(Client, ctx)
	if ConErr != nil {
		count += 1
		if count < 5 {
			log.Println("Reconnecting Database...")
		} else {
			log.Println("Could not connect to the Database...")
			goto label
		}
	}
	defer Client.Disconnect(ctx)

	fmt.Printf("%T Connected ...\n", Client)

}
