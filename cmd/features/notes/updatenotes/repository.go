package updatenotes

import (
	"context"
	"test/cmd/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNoteByID(ctx context.Context, conn *mongo.Database, id primitive.ObjectID) (*model.Notes, error) {
	var note *model.Notes

	filter := bson.M{
		"_id": id,
	}

	err := conn.Collection(model.NOTES_COLLECTION).FindOne(ctx, filter, nil).Decode(&note)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func updateNote(ctx context.Context, conn *mongo.Database, data *model.Notes) error {
	filter := bson.M{
		"_id": data.ID,
	}

	request := bson.M{
		"$set": bson.M{"text": data.Text},
	}

	_, err := conn.Collection(model.NOTES_COLLECTION).UpdateOne(ctx, filter, request, nil)
	if err != nil {
		return err
	}

	return nil
}
