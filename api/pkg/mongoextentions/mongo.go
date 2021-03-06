package mongohelper

import (
	"context"
	"github.com/cemilcan98/chess/pkg/log"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewDatabase(uri, databaseName string) (db *mongo.Database, err error) {

	log.Logger.Infof("Mongo:Connection Uri:%s", uri)
	clientOptions := options.
		Client().
		ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {

		return db, errors.Wrap(ConnectionError, err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Logger.Errorf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db, errors.Wrap(ConnectionError, err.Error())
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return db, errors.Wrap(PingError, err.Error())
	}

	db = client.Database(databaseName)

	return db, err
}
func BuildQuery(params map[string]string) (query bson.M) {

	query = bson.M{}
	for field, value := range params {

		if len(value) == 0 {
			continue
		}
		//if field == "_id" {
		//	objectIDS, _ := primitive.ObjectIDFromHex(value)
		//	query["_id"] = objectIDS
		//} else
		{
			query[field] = value
		}
	}
	return query
}
