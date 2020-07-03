package modeltests

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dzahariev/e2e-rest/api/controller"
	"github.com/dzahariev/e2e-rest/api/model"
	. "github.com/dzahariev/e2e-rest/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

var _ = Describe("Tests configuration", func() {
	var (
		server          = controller.Server{}
		dbName          = fmt.Sprintf("mod_%s", strings.ReplaceAll(GetID().String(), "-", ""))
		user1ID         = GetID()
		user2ID         = GetID()
		event1ID        = GetID()
		event2ID        = GetID()
		session1ID      = GetID()
		session2ID      = GetID()
		subscription1ID = GetID()
		subscription2ID = GetID()
		comment1ID      = GetID()
		comment2ID      = GetID()
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

	// Event
	eventEntityType := EntityType{
		Name:   "Event",
		Entity: &model.Event{},
		NewEntity: &model.Event{
			Base: model.Base{
				ID: event1ID,
			},
			Name: "Winter Summit",
			Year: "2020",
		},
		NewEntity1: &model.Event{
			Base: model.Base{
				ID: event2ID,
			},
			Name: "Summer Summit",
			Year: "2020",
		},
	}

	// Session
	sessionEntityType := EntityType{
		Name:   "Session",
		Entity: &model.Session{},
		NewEntity: &model.Session{
			Base: model.Base{
				ID: session1ID,
			},
			Name: "Main theme",
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Event: model.Event{
				Base: model.Base{
					ID: event1ID,
				},
				Name: "Winter Summit",
				Year: "2020",
			},
			EventID: event1ID,
		},
		NewEntity1: &model.Session{
			Base: model.Base{
				ID: session2ID,
			},
			Name: "Main theme 2",
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Event: model.Event{
				Base: model.Base{
					ID: event1ID,
				},
				Name: "Winter Summit",
				Year: "2020",
			},
			EventID: event1ID},
	}

	// Subscription
	subscriptionEntityType := EntityType{
		Name:   "Subscription",
		Entity: &model.Subscription{},
		NewEntity: &model.Subscription{
			Base: model.Base{
				ID: subscription1ID,
			},
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Session: model.Session{
				Base: model.Base{
					ID: session1ID,
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: user1ID,
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: user1ID,
				Event: model.Event{
					Base: model.Base{
						ID: event1ID,
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: event1ID,
			},
			SessionID: session1ID,
		},
		NewEntity1: &model.Subscription{
			Base: model.Base{
				ID: subscription2ID,
			},
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Session: model.Session{
				Base: model.Base{
					ID: session1ID,
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: user1ID,
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: user1ID,
				Event: model.Event{
					Base: model.Base{
						ID: event1ID,
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: event1ID,
			},
			SessionID: session1ID,
		},
	}

	// Comment
	commentEntityType := EntityType{
		Name:   "Comment",
		Entity: &model.Comment{},
		NewEntity: &model.Comment{
			Base: model.Base{
				ID: comment1ID,
			},
			Message: "Special comment!",
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Session: model.Session{
				Base: model.Base{
					ID: session1ID,
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: user1ID,
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: user1ID,
				Event: model.Event{
					Base: model.Base{
						ID: event1ID,
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: event1ID,
			},
			SessionID: session1ID,
		},
		NewEntity1: &model.Comment{
			Base: model.Base{
				ID: comment2ID,
			},
			Message: "Special comment!",
			User: model.User{
				Base: model.Base{
					ID: user1ID,
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: user1ID,
			Session: model.Session{
				Base: model.Base{
					ID: session1ID,
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: user1ID,
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: user1ID,
				Event: model.Event{
					Base: model.Base{
						ID: event1ID,
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: event1ID,
			},
			SessionID: session1ID,
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
		Entry(fmt.Sprintf("should successfully create new %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Fetch entity",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			err = entityType.Entity.FindByID(server.DB, entityType.NewEntity.GetID())
			Expect(err).ShouldNot(HaveOccurred())

			Expect(entityType.NewEntity.GetID()).To(BeEquivalentTo(entityType.Entity.GetID()))
		},
		Entry(fmt.Sprintf("should successfully fetch the %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully fetch the %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully fetch the %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully fetch the %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully fetch the %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Fetch all entities",
		func(entityType EntityType) {
			entitiesStart, err := entityType.Entity.FindAll(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			err = entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			err = entityType.NewEntity1.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			entities, err := entityType.Entity.FindAll(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(len(*entities)).To(BeEquivalentTo(len(*entitiesStart) + 2))
		},
		Entry(fmt.Sprintf("should successfully fetch all %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully fetch all %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully fetch all %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully fetch all %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully fetch all %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Update entity",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			now := time.Now()
			entityType.NewEntity.SetCreatedAt(now)
			err = entityType.NewEntity.Update(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(entityType.NewEntity.GetCreatedAt()).To(BeEquivalentTo(now))
		},
		Entry(fmt.Sprintf("should successfully update the %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully update the %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully update the %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully update the %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully update the %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Delete entity",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			err = entityType.NewEntity.Delete(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			countEnd, err := entityType.Entity.Count(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(countEnd).To(BeEquivalentTo(0))
		},
		Entry(fmt.Sprintf("should successfully delete the %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully delete the %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully delete the %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully delete the %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully delete the %s", commentEntityType.Name), commentEntityType),
	)
})
