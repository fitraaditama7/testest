package updatenotes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/pkg/responses"
	"test/pkg/router"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	dbConn *mongo.Database
}

func New(dbConn *mongo.Database) *handler {
	return &handler{dbConn}
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var body request
	var id = router.Param(ctx, "id")

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	fmt.Println(id)
	fmt.Println(id)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		responses.Error(w, errInvalidIDFormatError)
		return
	}

	note, err := getNoteByID(ctx, h.dbConn, objID)
	if err == mongo.ErrNoDocuments {
		log.Println(err)
		responses.Error(w, errNoteNotFoundError)
		return
	}
	if err != nil {
		log.Println(err)
		responses.Error(w, errUnexpectedError)
		return
	}

	note.Text = body.Text
	err = updateNote(ctx, h.dbConn, note)
	if err != nil {
		log.Println(err)
		responses.Error(w, errUnexpectedError)
		return
	}

	responses.Success(w, "SUCCESS")
}
