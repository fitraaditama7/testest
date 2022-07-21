package insertnotes

import (
	"encoding/json"
	"log"
	"net/http"
	"test/pkg/responses"

	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	dbConn *mongo.Database
}

func New(dbConn *mongo.Database) *handler {
	return &handler{dbConn}
}

func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var body request

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	err := insertNote(ctx, h.dbConn, &body)
	if err != nil {
		log.Println(err)
		responses.Error(w, errUnexpectedError)
		return
	}

	responses.Success(w, "SUCCESS")
}
