package insertnotes

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"test/cmd/domain/model"
	"test/pkg/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoUri string

var _ = Describe("Handler", func() {
	Describe("ListNotes", func() {
		var db *mongo.Database
		var notes = model.NotesMock()

		BeforeEach(func() {
			var conf = database.DBConfigTest
			conf.Host = mongoUri

			conn, err := database.InitMongoDB(conf)
			Expect(err).NotTo(HaveOccurred())

			db = conn.Database(conf.Name)
			Expect(database.TruncateDataForTest(conn.Database(conf.Name))).To(Succeed())
		})

		tests := []struct {
			name           string
			expectedStatus int
			expectedNotes  *model.Notes
		}{
			{
				name:           "be able to insert notes",
				expectedStatus: http.StatusOK,
				expectedNotes:  notes,
			},
		}

		for _, test := range tests {
			test := test
			It(test.name, func() {
				ctx := context.Background()
				b, err := json.Marshal(test.expectedNotes)
				Expect(err).NotTo(HaveOccurred())

				handler := New(db)

				req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(string(b)))
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				handler.Insert(w, req)

				if w.Code == http.StatusOK {
					var note *model.Notes
					filter := bson.M{
						"text": test.expectedNotes.Text,
					}

					err = db.Collection(model.NOTES_COLLECTION).FindOne(ctx, filter, nil).Decode(&note)
					Expect(err).NotTo(HaveOccurred())

					Expect(test.expectedNotes.Text).To(Equal(note.Text))
				}
				Expect(w).To(HaveHTTPStatus(test.expectedStatus))
			})
		}
	})
})
