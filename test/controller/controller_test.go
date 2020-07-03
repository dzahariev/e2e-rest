package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

var (
	server       = controller.Server{}
	dbName       = fmt.Sprintf("ctr_%s", strings.ReplaceAll(GetID().String(), "-", ""))
	logedUserID  = GetID()
	userPassword = "secret007"

	loggedUser = model.User{
		Base: model.Base{
			ID: logedUserID,
		},
		Name:     "Super Admin",
		Email:    "super.admin@mymail.local",
		Password: userPassword,
	}

	validLoginPayload = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    loggedUser.Email,
		Password: userPassword,
	}
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

func CreateUserAndGetToken(server *controller.Server) string {
	loggedUser.Password = userPassword
	err := loggedUser.Save(server.DB)
	Expect(err).ShouldNot(HaveOccurred())

	token, err := server.GetTokenForUser(validLoginPayload.Email, validLoginPayload.Password)
	Expect(err).ShouldNot(HaveOccurred())
	return fmt.Sprintf("Bearer %v", token)
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
		server.Initialize(dbUser, dbPassword, dbPort, dbHost, dbName)
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

	DescribeTable("Get all for entity should return OK with valid token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			request, err := http.NewRequest("GET", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", token)

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusOK))
		},
		Entry(fmt.Sprintf("should successfully get all %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully get all %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully get all %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully get all %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully get all %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Get all for entity should return Status Unauthorized with wrong token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)
			request, err := http.NewRequest("GET", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", fmt.Sprintf("%sinv", token))

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to get all  %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Get all for entity should return Status Unauthorized when token is not provided",
		func(entityType EntityType) {
			request, err := http.NewRequest("GET", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), nil)
			Expect(err).ShouldNot(HaveOccurred())

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to get all %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to get all %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Create entity should return OK with valid token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)
			entityJSON, err := json.Marshal(entityType.NewEntity)
			Expect(err).ShouldNot(HaveOccurred())
			request, err := http.NewRequest("POST", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", token)

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusCreated))
		},
		Entry(fmt.Sprintf("should successfully fetch %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully fetch %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully create new %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Create entity should return Status Unauthorized with wrong token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)
			entityJSON, err := json.Marshal(entityType.NewEntity)
			Expect(err).ShouldNot(HaveOccurred())
			request, err := http.NewRequest("POST", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", fmt.Sprintf("%sinv", token))

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		//Entry(fmt.Sprintf("should fail to create new %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Create entity should return Status Unauthorized when token is not provided",
		func(entityType EntityType) {
			entityJSON, err := json.Marshal(entityType.NewEntity)
			Expect(err).ShouldNot(HaveOccurred())
			request, err := http.NewRequest("POST", fmt.Sprintf("/%s", strings.ToLower(entityType.Name)), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		//Entry(fmt.Sprintf("should fail to create new %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to create new %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Get single entity should return OK with valid token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("GET", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", token)

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusOK))
		},
		Entry(fmt.Sprintf("should successfully get single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully get single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully get single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully get single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully get single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Get single entity should return Status Unauthorized with wrong token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("GET", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", fmt.Sprintf("%sinv", token))

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to get single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Get single entity should return Status Unauthorized when token is not provided",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("GET", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to get single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to get single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Update single entity should return OK with valid token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			entityType.NewEntity.SetCreatedAt(time.Now())

			entityJSON, err := json.Marshal(entityType.NewEntity1)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("PUT", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", token)

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusOK))
		},
		// Entry(fmt.Sprintf("should successfully update single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully update single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully update single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully update single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully update single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Update single entity should return Status Unauthorized with wrong token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			entityType.NewEntity.SetCreatedAt(time.Now())

			entityJSON, err := json.Marshal(entityType.NewEntity1)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("PUT", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", fmt.Sprintf("%sinv", token))

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to update single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Update single entity should return Status Unauthorized when token is not provided",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())
			entityType.NewEntity.SetCreatedAt(time.Now())

			entityJSON, err := json.Marshal(entityType.NewEntity1)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("PUT", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), bytes.NewBufferString(string(entityJSON)))
			Expect(err).ShouldNot(HaveOccurred())

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		Entry(fmt.Sprintf("should fail to update single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to update single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Delete single entity should return NoContent with valid token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", token)

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusNoContent))
		},
		// Entry(fmt.Sprintf("should successfully delete single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should successfully delete single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should successfully delete single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should successfully delete single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should successfully delete single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Delete single entity should return Status Unauthorized with wrong token",
		func(entityType EntityType) {
			token := CreateUserAndGetToken(&server)

			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())
			request.Header.Set("Authorization", fmt.Sprintf("%sinv", token))

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		// Entry(fmt.Sprintf("should fail to delete single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", commentEntityType.Name), commentEntityType),
	)

	DescribeTable("Delete single entity should return Status Unauthorized when token is not provided",
		func(entityType EntityType) {
			err := entityType.NewEntity.Save(server.DB)
			Expect(err).ShouldNot(HaveOccurred())

			request, err := http.NewRequest("DELETE", fmt.Sprintf("/%s/%s", strings.ToLower(entityType.Name), entityType.NewEntity.GetID().String()), nil)
			Expect(err).ShouldNot(HaveOccurred())

			requestRecorder := httptest.NewRecorder()
			server.Router.ServeHTTP(requestRecorder, request)
			Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
		},
		// Entry(fmt.Sprintf("should fail to delete single %s", userEntityType.Name), userEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", eventEntityType.Name), eventEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", sessionEntityType.Name), sessionEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", subscriptionEntityType.Name), subscriptionEntityType),
		Entry(fmt.Sprintf("should fail to delete single %s", commentEntityType.Name), commentEntityType),
	)

})
