package logintests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dzahariev/e2e-rest/api/controller"
	"github.com/dzahariev/e2e-rest/api/model"
	. "github.com/dzahariev/e2e-rest/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLogin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Login Suite")
}

var _ = Describe("Tests configuration", func() {
	var (
		server        = controller.Server{}
		dbName        = fmt.Sprintf("lgn_%s", strings.ReplaceAll(GetID().String(), "-", ""))
		validPassword = "secret007"
		user1ID       = GetID()
		user          = model.User{}
	)

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

		// Init vars
		// User
		user = model.User{
			Base: model.Base{
				ID: user1ID,
			},
			Name:     "Joe Satriani",
			Email:    "joe.satriani@mymail.local",
			Password: validPassword,
		}

	})

	var _ = Describe("Login test", func() {
		Context("check the password hash ", func() {
			It(fmt.Sprintf("should be the same as saved for user %s", user.Name), func() {
				err := user.Save(server.DB)
				Expect(err).ShouldNot(HaveOccurred())

				_, err = server.GetTokenForUser(user.Email, validPassword)
				Expect(err).ShouldNot(HaveOccurred())
				_, err = server.GetTokenForUser(user.Email, fmt.Sprintf("inv%s", validPassword))
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("check the login API ", func() {
			It(fmt.Sprintf("should login with valid user and password"), func() {
				err := user.Save(server.DB)
				Expect(err).ShouldNot(HaveOccurred())
				newUser := struct {
					Email    string `json:"email"`
					Password string `json:"password"`
				}{
					Email:    user.Email,
					Password: validPassword,
				}

				userJSON, err := json.Marshal(newUser)
				Expect(err).ShouldNot(HaveOccurred())
				request, err := http.NewRequest("POST", "/login", bytes.NewBufferString(string(userJSON)))
				Expect(err).ShouldNot(HaveOccurred())

				requestRecorder := httptest.NewRecorder()
				handler := http.HandlerFunc(server.LogIn)
				handler.ServeHTTP(requestRecorder, request)
				Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusOK))
			})

			It(fmt.Sprintf("should not login with invalid user name"), func() {
				err := user.Save(server.DB)
				Expect(err).ShouldNot(HaveOccurred())
				newUser := struct {
					Email    string `json:"email"`
					Password string `json:"password"`
				}{
					Email:    fmt.Sprintf("inv%s", user.Email),
					Password: validPassword,
				}

				userJSON, err := json.Marshal(newUser)
				Expect(err).ShouldNot(HaveOccurred())
				request, err := http.NewRequest("POST", "/login", bytes.NewBufferString(string(userJSON)))
				Expect(err).ShouldNot(HaveOccurred())

				requestRecorder := httptest.NewRecorder()
				handler := http.HandlerFunc(server.LogIn)
				handler.ServeHTTP(requestRecorder, request)
				Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnprocessableEntity))
			})

			It(fmt.Sprintf("should not login with invalid password"), func() {
				err := user.Save(server.DB)
				Expect(err).ShouldNot(HaveOccurred())
				newUser := struct {
					Email    string `json:"email"`
					Password string `json:"password"`
				}{
					Email:    user.Email,
					Password: fmt.Sprintf("inv%s", validPassword),
				}

				userJSON, err := json.Marshal(newUser)
				Expect(err).ShouldNot(HaveOccurred())
				request, err := http.NewRequest("POST", "/login", bytes.NewBufferString(string(userJSON)))
				Expect(err).ShouldNot(HaveOccurred())

				requestRecorder := httptest.NewRecorder()
				handler := http.HandlerFunc(server.LogIn)
				handler.ServeHTTP(requestRecorder, request)
				Expect(requestRecorder.Code).Should(BeEquivalentTo(http.StatusUnauthorized))
			})

		})

	})
})
