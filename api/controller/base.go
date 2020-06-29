package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Server represent current API server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// DBInitialize is used to init a DB cnnection
func (server *Server) DBInitialize(dbUser, dbPassword, dbPort, dbHost, dbName string) {
	dbDriver := "postgres"
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	var err error
	server.DB, err = gorm.Open(dbDriver, DBURL)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot connect to %s database with error: %v", dbDriver, err))
	}
	log.Printf("We are connected to the %s database", dbDriver)

	server.DB.AutoMigrate(&model.User{}, &model.Event{}, &model.Session{}, &model.Subscription{}, &model.Comment{})
}

// RoutesInitialize is used to register routes
func (server *Server) RoutesInitialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Initialize is used to init a DB cnnection and register routes
func (server *Server) Initialize(dbUser, dbPassword, dbPort, dbHost, dbName string) {
	server.DBInitialize(dbUser, dbPassword, dbPort, dbHost, dbName)
	server.RoutesInitialize()
}

// Run starts the http server
func (server *Server) Run(addr string) {
	log.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
