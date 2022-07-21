package listnotes

import (
	"context"
	"net/http"
	"net/http/httptest"
	"test/cmd/domain/model"
	"test/pkg/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
				name:           "be able to get notes",
				expectedStatus: http.StatusOK,
				expectedNotes:  notes,
			},
		}

		for _, test := range tests {
			test := test
			It(test.name, func() {
				ctx := context.Background()

				_, err := db.Collection(model.NOTES_COLLECTION).InsertOne(ctx, test.expectedNotes, nil)
				Expect(err).NotTo(HaveOccurred())

				handler := New(db)

				req, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				handler.List(w, req)

				Expect(w).To(HaveHTTPStatus(test.expectedStatus))
			})
		}
	})
})
