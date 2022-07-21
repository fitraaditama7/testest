package listnotes

import (
	"context"
	"test/pkg/database"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ListNotes")
}

var container testcontainers.Container
var containerCtx = context.Background()

var _ = BeforeSuite(func() {
	var err error
	containers, err := database.StartMongoContainer(containerCtx)
	Expect(err).NotTo(HaveOccurred())
	container = containers.Container
	mongoUri = containers.URI
})

var _ = AfterSuite(func() {
	_ = container.Terminate(containerCtx)
})
