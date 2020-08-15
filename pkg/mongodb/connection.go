package mongodb

import (
	"context"
	"time"

	"github.com/darksasori/finance/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var conn *mongo.Database

func Connect(ctx context.Context) error {
	uri := utils.GetEnv("MONGODB_URI", "mongodb://localhost:27017")
	db := utils.GetEnv("MONGODB_DB", "testing")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	database := client.Database(db)
	conn = database

	return err
}
