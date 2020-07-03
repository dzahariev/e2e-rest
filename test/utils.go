package test

import (
	"fmt"
	"log"
	"os"

	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// EntityType represent abstraction for the diferent entities
type EntityType struct {
	Name       string
	Entity     model.Object
	NewEntity  model.Object
	NewEntity1 model.Object
}

// CreateDB creates the database
func CreateDB(DB *gorm.DB, dbUser, dbPassword, dbPort, dbHost, dbName string) error {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, "postgres", dbPassword)
	DB, err := gorm.Open("postgres", DBURL)
	defer DB.Close()
	if err != nil {
		log.Printf("Unable to conect DB %s", dbName)
		return err
	}
	statement := fmt.Sprintf("CREATE DATABASE %s;", dbName)
	DB = DB.Exec(statement)
	if DB.Error != nil {
		log.Printf("Unable to create DB %s with error %v", dbName, err)
		return err
	}
	return nil
}

// DropDB drops the database
func DropDB(DB *gorm.DB, dbUser, dbPassword, dbPort, dbHost, dbName string) error {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, "postgres", dbPassword)
	DB, err := gorm.Open("postgres", DBURL)
	defer DB.Close()
	if err != nil {
		log.Printf("Unable to conect DB %s	", dbName)
		return err
	}
	statement := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	DB = DB.Exec(statement)
	if DB.Error != nil {
		log.Printf("Unable to drop the DB %s", dbName)
		return err
	}
	return nil
}

// RecreateTables recreates the tables in database
func RecreateTables(DB *gorm.DB) error {
	err := DB.DropTableIfExists(&model.User{}).Error
	if err != nil {
		return err
	}
	err = DB.DropTableIfExists(&model.Event{}).Error
	if err != nil {
		return err
	}
	err = DB.DropTableIfExists(&model.Session{}).Error
	if err != nil {
		return err
	}
	err = DB.DropTableIfExists(&model.Subscription{}).Error
	if err != nil {
		return err
	}
	err = DB.DropTableIfExists(&model.Comment{}).Error
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&model.User{}).Error
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Event{}).Error
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Session{}).Error
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Subscription{}).Error
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// LoadEnvironment for testing
func LoadEnvironment() error {
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		return err
	}
	return nil
}

// GetID returns an ID
func GetID() uuid.UUID {
	ID, _ := uuid.NewV4()
	return ID
}
