package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mdbOptions *options.ClientOptions
	mdbConn    string
)

func InitMongoDBSettings(connstring string) {
	mdbConn = connstring
	mdbOptions = options.Client().ApplyURI(connstring).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
}

type MongoRequest func(*mongo.Client) error

// execute mongodb functions with specific context
func RunMDBWithContext(ctx context.Context, f MongoRequest) (err error) {
	// create a new client which will be used for this request
	var client *mongo.Client
	if client, err = mongo.Connect(ctx, mdbOptions); err == nil {
		if err = f(client); err == nil {
			return client.Disconnect(context.TODO())
		}
	}
	return
}

// execute mongodb functions with default context
func RunMDB(f MongoRequest) error {
	return RunMDBWithContext(context.TODO(), f)
}
