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

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")

	server.Initialize(dbUser, dbPassword, dbPort, dbHost, dbName)
	server.Run(":8080")
}
