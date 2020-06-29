package test

import (
	"log"
	"os"

	"github.com/dzahariev/e2e-rest/api/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const (
	idUser          = "ef2c105a-93fc-4599-bcfa-a10fff709100"
	idUser2         = "ef2c105a-93fc-4599-bcfa-a10fff709101"
	idEvent         = "ef2c105a-93fc-4599-bcfa-a10fff709102"
	idEvent2        = "ef2c105a-93fc-4599-bcfa-a10fff709103"
	idSession       = "ef2c105a-93fc-4599-bcfa-a10fff709104"
	idSession2      = "ef2c105a-93fc-4599-bcfa-a10fff709105"
	idSubscription  = "ef2c105a-93fc-4599-bcfa-a10fff709106"
	idSubscription2 = "ef2c105a-93fc-4599-bcfa-a10fff709107"
	idComment       = "ef2c105a-93fc-4599-bcfa-a10fff709108"
	idComment2      = "ef2c105a-93fc-4599-bcfa-a10fff709109"
)

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
	log.Printf("All tables recreated!")
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

// GetUserID returns an ID
func GetUserID() uuid.UUID {
	ID, _ := uuid.FromString(idUser)
	return ID
}

// GetUserID2 returns an ID
func GetUserID2() uuid.UUID {
	ID, _ := uuid.FromString(idUser2)
	return ID
}

// GetEventID returns an ID
func GetEventID() uuid.UUID {
	ID, _ := uuid.FromString(idEvent)
	return ID
}

// GetEventID2 returns an ID
func GetEventID2() uuid.UUID {
	ID, _ := uuid.FromString(idEvent2)
	return ID
}

// GetSessionID returns an ID
func GetSessionID() uuid.UUID {
	ID, _ := uuid.FromString(idSession)
	return ID
}

// GetSessionID2 returns an ID
func GetSessionID2() uuid.UUID {
	ID, _ := uuid.FromString(idSession2)
	return ID
}

// GetSubscriptionID returns an ID
func GetSubscriptionID() uuid.UUID {
	ID, _ := uuid.FromString(idSubscription)
	return ID
}

// GetSubscriptionID2 returns an ID
func GetSubscriptionID2() uuid.UUID {
	ID, _ := uuid.FromString(idSubscription2)
	return ID
}

// GetCommentID returns an ID
func GetCommentID() uuid.UUID {
	ID, _ := uuid.FromString(idComment)
	return ID
}

// GetCommentID2 returns an ID
func GetCommentID2() uuid.UUID {
	ID, _ := uuid.FromString(idComment2)
	return ID
}
