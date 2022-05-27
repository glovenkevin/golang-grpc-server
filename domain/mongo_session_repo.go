package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoSessionRepo interface {
	CreateSession() (mongo.Session, error)
	InsertLogs(ctx context.Context, payload interface{}) error
}
