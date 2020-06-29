package main

import (
	"log"
	"os"

	"github.com/dzahariev/e2e-rest/api/controller"
	"github.com/joho/godotenv"
)

var server = controller.Server{}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not loaded due to:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	server.Initialize(dbUser, dbPassword, dbPort, dbHost, dbName)
	server.Run(":8080")
}
