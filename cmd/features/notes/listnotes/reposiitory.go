package listnotes

import (
	"context"
	"test/cmd/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNotes(ctx context.Context, conn *mongo.Database) ([]*model.Notes, error) {
	var notes []*model.Notes

	filter := bson.M{}
	curr, err := conn.Collection(model.NOTES_COLLECTION).Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	defer curr.Close(ctx)

	for curr.Next(ctx) {
		var note = new(model.Notes)
		if err := curr.Decode(&note); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}
