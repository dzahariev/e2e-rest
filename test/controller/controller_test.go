package controllertests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dzahariev/e2e-rest/api/controller"
	"github.com/dzahariev/e2e-rest/api/model"
	. "github.com/dzahariev/e2e-rest/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = Describe("Tests configuration", func() {
	var (
		server  = controller.Server{}
		dbName  = fmt.Sprintf("ctr_%s", strings.ReplaceAll(GetID().String(), "-", ""))
		user1ID = GetID()
		user2ID = GetID()
	)
	// User
	userEntityType := EntityType{
		Name:   "User",
		Entity: &model.User{},
		NewEntity: &model.User{
			Base: model.Base{
				ID: user1ID,
			},
			Name:     "Joe Satriani",
			Email:    "joe.satriani@mymail.local",
			Password: "secret007",
		},
		NewEntity1: &model.User{
			Base: model.Base{
				ID: user2ID,
			},
			Name:     "John Smith",
			Email:    "john.smith@mymail.local",
			Password: "secret007",
		},
	}

	BeforeSuite(func() {
		err := LoadEnvironment()
		Expect(err).ShouldNot(HaveOccurred())

		dbUser := os.Getenv("TEST_POSTGRES_USER")
		dbPassword := os.Getenv("TEST_POSTGRES_PASSWORD")
		dbPort := os.Getenv("TEST_POSTGRES_PORT")
		dbHost := os.Getenv("TEST_POSTGRES_HOST")

		err = CreateDB(server.DB, dbUser, dbPassword, dbPort, dbHost, dbName)
		Expect(err).ShouldNot(HaveOccurred())
		server.DBInitialize(dbUser, dbPassword, dbPort, dbHost, dbName)
	})

	AfterSuite(func() {
		dbUser := os.Getenv("TEST_POSTGRES_USER")
		dbPassword := os.Getenv("TEST_POSTGRES_PASSWORD")
		dbPort := os.Getenv("TEST_POSTGRES_PORT")
		dbHost := os.Getenv("TEST_POSTGRES_HOST")

		server.DB.Close()
		err := DropDB(server.DB, dbUser, dbPassword, dbPort, dbHost, dbName)
		Expect(err).ShouldNot(HaveOccurred())
	})

	BeforeEach(func() {
		err := RecreateTables(server.DB)
		Expect(err).ShouldNot(HaveOccurred())
	})

	DescribeTable("New entity",
		func(entityType EntityType) {
			countStart, err := entityType.Entity.Count(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			err = entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			countEnd, err := entityType.Entity.Count(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(countEnd - countStart).To(BeEquivalentTo(1))
		},
		Entry(fmt.Sprintf("should successfully create new %s", userEntityType.Name), userEntityType),
	)
})
