package logging

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	cli        *mongo.Client
	dbName     string
	collection string
}

func NewLogRepository(client *mongo.Client, db, collect string) *Repository {
	return &Repository{
		cli:        client,
		dbName:     db,
		collection: collect,
	}
}

func (r *Repository) Insert(item PingLogMessage) error {
	_, err := r.cli.Database(r.dbName).
		Collection(r.collection).
		InsertOne(context.TODO(), item)
	if err != nil {
		return fmt.Errorf("insert log item err: %w", err)
	}
	return nil
}
