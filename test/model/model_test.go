package modeltests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dzahariev/e2e-rest/api/controller"
	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/dzahariev/e2e-rest/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

type EntityType struct {
	Name       string
	Entity     model.Object
	NewEntity  model.Object
	NewEntity1 model.Object
}

var _ = Describe("Tests configuration", func() {
	var server = controller.Server{}

	// User
	userEntityType := EntityType{
		Name:   "User",
		Entity: &model.User{},
		NewEntity: &model.User{
			Base: model.Base{
				ID: test.GetUserID(),
			},
			Name:     "Joe Satriani",
			Email:    "joe.satriani@mymail.local",
			Password: "secret007",
		},
		NewEntity1: &model.User{
			Base: model.Base{
				ID: test.GetUserID2(),
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
				ID: test.GetEventID(),
			},
			Name: "Winter Summit",
			Year: "2020",
		},
		NewEntity1: &model.Event{
			Base: model.Base{
				ID: test.GetEventID2(),
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
				ID: test.GetSessionID(),
			},
			Name: "Main theme",
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Event: model.Event{
				Base: model.Base{
					ID: test.GetEventID(),
				},
				Name: "Winter Summit",
				Year: "2020",
			},
			EventID: test.GetEventID(),
		},
		NewEntity1: &model.Session{
			Base: model.Base{
				ID: test.GetSessionID2(),
			},
			Name: "Main theme 2",
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Event: model.Event{
				Base: model.Base{
					ID: test.GetEventID(),
				},
				Name: "Winter Summit",
				Year: "2020",
			},
			EventID: test.GetEventID()},
	}

	// Subscription
	subscriptionEntityType := EntityType{
		Name:   "Subscription",
		Entity: &model.Subscription{},
		NewEntity: &model.Subscription{
			Base: model.Base{
				ID: test.GetSubscriptionID(),
			},
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Session: model.Session{
				Base: model.Base{
					ID: test.GetSessionID(),
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: test.GetUserID(),
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: test.GetUserID(),
				Event: model.Event{
					Base: model.Base{
						ID: test.GetEventID(),
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: test.GetEventID(),
			},
			SessionID: test.GetSessionID(),
		},
		NewEntity1: &model.Subscription{
			Base: model.Base{
				ID: test.GetSubscriptionID2(),
			},
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Session: model.Session{
				Base: model.Base{
					ID: test.GetSessionID(),
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: test.GetUserID(),
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: test.GetUserID(),
				Event: model.Event{
					Base: model.Base{
						ID: test.GetEventID(),
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: test.GetEventID(),
			},
			SessionID: test.GetSessionID(),
		},
	}

	// Comment
	commentEntityType := EntityType{
		Name:   "Comment",
		Entity: &model.Comment{},
		NewEntity: &model.Comment{
			Base: model.Base{
				ID: test.GetCommentID(),
			},
			Message: "Special comment!",
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Session: model.Session{
				Base: model.Base{
					ID: test.GetSessionID(),
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: test.GetUserID(),
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: test.GetUserID(),
				Event: model.Event{
					Base: model.Base{
						ID: test.GetEventID(),
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: test.GetEventID(),
			},
			SessionID: test.GetSessionID(),
		},
		NewEntity1: &model.Comment{
			Base: model.Base{
				ID: test.GetCommentID2(),
			},
			Message: "Special comment!",
			User: model.User{
				Base: model.Base{
					ID: test.GetUserID(),
				},
				Name:     "Joe Satriani",
				Email:    "joe.satriani@mymail.local",
				Password: "secret007",
			},
			UserID: test.GetUserID(),
			Session: model.Session{
				Base: model.Base{
					ID: test.GetSessionID(),
				},
				Name: "Main theme",
				User: model.User{
					Base: model.Base{
						ID: test.GetUserID(),
					},
					Name:     "Joe Satriani",
					Email:    "joe.satriani@mymail.local",
					Password: "secret007",
				},
				UserID: test.GetUserID(),
				Event: model.Event{
					Base: model.Base{
						ID: test.GetEventID(),
					},
					Name: "Winter Summit",
					Year: "2020",
				},
				EventID: test.GetEventID(),
			},
			SessionID: test.GetSessionID(),
		},
	}

	BeforeSuite(func() {
		err := test.LoadEnvironment()
		Expect(err).ShouldNot(HaveOccurred())

		dbUser := os.Getenv("TEST_POSTGRES_USER")
		dbPassword := os.Getenv("TEST_POSTGRES_PASSWORD")
		dbPort := os.Getenv("TEST_POSTGRES_PORT")
		dbHost := os.Getenv("TEST_POSTGRES_HOST")
		dbName := os.Getenv("TEST_POSTGRES_DB")

		server.DBInitialize(dbUser, dbPassword, dbPort, dbHost, dbName)
	})

	BeforeEach(func() {
		err := test.RecreateTables(server.DB)
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
