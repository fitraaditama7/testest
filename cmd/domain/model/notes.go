package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const NOTES_COLLECTION = "notes"

type Notes struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Text string             `json:"text" bson:"text"`
}

func NotesMock() *Notes {
	return &Notes{
		ID:   primitive.NewObjectID(),
		Text: "test",
	}
}
