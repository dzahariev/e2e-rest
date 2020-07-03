package controller

import "github.com/dzahariev/e2e-rest/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.ContentTypeJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middleware.ContentTypeJSON(s.LogIn)).Methods("POST")

	// User routes
	s.Router.HandleFunc("/user", middleware.ContentTypeJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteUser))).Methods("DELETE")

	// Event routes
	s.Router.HandleFunc("/event", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.CreateEvent))).Methods("POST")
	s.Router.HandleFunc("/event", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetEvents))).Methods("GET")
	s.Router.HandleFunc("/event/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetEvent))).Methods("GET")
	s.Router.HandleFunc("/event/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateEvent))).Methods("PUT")
	s.Router.HandleFunc("/event/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteEvent))).Methods("DELETE")

	// Session routes
	s.Router.HandleFunc("/session", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.CreateSession))).Methods("POST")
	s.Router.HandleFunc("/session", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetSessions))).Methods("GET")
	s.Router.HandleFunc("/session/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetSession))).Methods("GET")
	s.Router.HandleFunc("/session/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateSession))).Methods("PUT")
	s.Router.HandleFunc("/session/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteSession))).Methods("DELETE")

	// Subscription routes
	s.Router.HandleFunc("/subscription", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.CreateSubscription))).Methods("POST")
	s.Router.HandleFunc("/subscription", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetSubscriptions))).Methods("GET")
	s.Router.HandleFunc("/subscription/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetSubscription))).Methods("GET")
	s.Router.HandleFunc("/subscription/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateSubscription))).Methods("PUT")
	s.Router.HandleFunc("/subscription/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteSubscription))).Methods("DELETE")

	// Comment routes
	s.Router.HandleFunc("/comment", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.CreateComment))).Methods("POST")
	s.Router.HandleFunc("/comment", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetComments))).Methods("GET")
	s.Router.HandleFunc("/comment/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.GetComment))).Methods("GET")
	s.Router.HandleFunc("/comment/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.UpdateComment))).Methods("PUT")
	s.Router.HandleFunc("/comment/{id}", middleware.ContentTypeJSON(middleware.CheckAuthentication(s.DeleteComment))).Methods("DELETE")

}
