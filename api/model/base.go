package model

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Object is an abstration of all Base objects
type Object interface {
	GetID() uuid.UUID
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	Save(db *gorm.DB) error
	Count(db *gorm.DB) (int, error)
	FindAll(db *gorm.DB) (*[]Object, error)
	FindByID(db *gorm.DB, uid uuid.UUID) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
}

// Base holds technical fields
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare initilises techncal fields
func (b *Base) Prepare() error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	b.ID = uuid
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now

	return nil
}
