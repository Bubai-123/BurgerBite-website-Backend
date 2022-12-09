package mapper

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllUser(client *mongo.Client) *mongo.Collection {
	usersCollection := client.Database("test").Collection("user")

	return usersCollection
}
