package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoSessionRepo struct {
	collection string
	db         *mongo.Database
	log        *log.Logger
}

func NewMonogoSessionRepo(db *mongo.Database, log *log.Logger) *MongoSessionRepo {
	return &MongoSessionRepo{collection: "mongo_session", db: db, log: log}
}

func (c *MongoSessionRepo) CreateSession() (mongo.Session, error) {
	ss, err := c.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	err = ss.StartTransaction()
	if err != nil {
		return nil, err
	}

	return ss, nil
}

func (c *MongoSessionRepo) InsertLogs(ctx context.Context, payload interface{}) error {
	collection := c.db.Collection(c.collection)
	_, err := collection.InsertOne(ctx, payload)
	return err
}
