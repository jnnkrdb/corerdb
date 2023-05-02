package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mdbClient *mongo.Client
)

func InitMongoDBSettings(connstring string) (err error) {
	if mdbClient, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(connstring).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))); err == nil {
		return err
	}
	return nil
}

// close the mongoedb connection
func CloseMDBConnection() error {
	return mdbClient.Disconnect(context.TODO())
}

// execute mongodb functions with default context
func RunMDB(f func(*mongo.Client) error) error {
	return f(mdbClient)
}
