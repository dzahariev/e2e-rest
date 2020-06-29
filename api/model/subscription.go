package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Subscription represents a session subscription
type Subscription struct {
	Base
	User      User `json:"user"`
	UserID    uuid.UUID
	Session   Session `json:"session"`
	SessionID uuid.UUID
}

// GetID returns the ID
func (s *Subscription) GetID() uuid.UUID {
	return s.ID
}

// GetCreatedAt returns the CreatedAt
func (s *Subscription) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetCreatedAt sets the CreatedAt
func (s *Subscription) SetCreatedAt(createdAt time.Time) {
	s.CreatedAt = createdAt
}

// Validate checks structure consistency
func (s *Subscription) Validate(action string) error {
	// always check
	if s.User.Name == "" {
		return fmt.Errorf("required User")
	}

	if s.Session.Name == "" {
		return fmt.Errorf("required Session")
	}

	return nil
}

// Save saves the structure as new object
func (s *Subscription) Save(db *gorm.DB) error {
	s.Prepare()

	err := s.Validate("update")
	if err != nil {
		return err
	}

	err = db.Create(&s).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll returns all known objects of this type
func (s *Subscription) FindAll(db *gorm.DB) (*[]Object, error) {
	entites := []Subscription{}
	err := db.Model(&s).Limit(100).Find(&entites).Error
	if err != nil {
		return &[]Object{}, err
	}

	objects := []Object{}
	for _, currentEntity := range entites {
		objects = append(objects, &currentEntity)
	}
	return &objects, nil
}

// Count returns count of all known objects of this type
func (s *Subscription) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&s).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByID returns an objects with corresponding ID if exists
func (s *Subscription) FindByID(db *gorm.DB, uid uuid.UUID) error {
	err := db.Model(&s).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates the existing objects
func (s *Subscription) Update(db *gorm.DB) error {
	if s.ID == uuid.Nil {
		return fmt.Errorf("cannot update non saved subscription")
	}

	err := s.Validate("update")
	if err != nil {
		return err
	}

	err = db.Model(&s).Updates(Subscription{
		User:    s.User,
		Session: s.Session,
		Base: Base{
			UpdatedAt: time.Now(),
		},
	}).Error

	if err != nil {
		return err
	}

	return nil
}

// Delete is removing existing objects
func (s *Subscription) Delete(db *gorm.DB) error {
	err := db.Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}
