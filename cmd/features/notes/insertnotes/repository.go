package insertnotes

import (
	"context"
	"test/cmd/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertNote(ctx context.Context, conn *mongo.Database, data *request) error {
	_, err := conn.Collection(model.NOTES_COLLECTION).InsertOne(ctx, bson.M{"text": data.Text}, nil)
	if err != nil {
		return err
	}

	return nil
}
