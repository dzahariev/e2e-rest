package controller

import "github.com/dzahariev/e2e-rest/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.ContentTypeJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middleware.ContentTypeJSON(s.LogIn)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middleware.ContentTypeJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteUser))).Methods("DELETE")
}
