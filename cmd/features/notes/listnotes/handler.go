package listnotes

import (
	"log"
	"net/http"
	"test/cmd/domain/model"
	"test/pkg/responses"

	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	dbConn *mongo.Database
}

func New(dbConn *mongo.Database) *handler {
	return &handler{dbConn}
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var notes []*model.Notes
	var err error

	notes, err = getNotes(ctx, h.dbConn)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println(err)
		responses.Error(w, errUnexpectedError)
		return
	}

	responses.Success(w, notes)
}
