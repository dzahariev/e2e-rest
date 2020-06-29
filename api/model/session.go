package model

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Session represents a session
type Session struct {
	Base
	Name          string `gorm:"size:255;not null;unique" json:"name"`
	User          User   `json:"author"`
	UserID        uuid.UUID
	Event         Event `json:"event"`
	EventID       uuid.UUID
	Subscriptions []Subscription `gorm:"foreignkey:SessionID"`
	Comments      []Comment      `gorm:"foreignkey:SessionID"`
}

// GetID returns the ID
func (s *Session) GetID() uuid.UUID {
	return s.ID
}

// GetCreatedAt returns the CreatedAt
func (s *Session) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetCreatedAt sets the CreatedAt
func (s *Session) SetCreatedAt(createdAt time.Time) {
	s.CreatedAt = createdAt
}

// Validate checks structure consistency
func (s *Session) Validate(action string) error {
	// always check
	if s.Name == "" {
		return fmt.Errorf("required Name")
	}

	if s.User.Name == "" {
		return fmt.Errorf("required Author")
	}

	if s.Event.Name == "" {
		return fmt.Errorf("required Event")
	}

	return nil
}

// Save saves the structure as new object
func (s *Session) Save(db *gorm.DB) error {
	s.Prepare()
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))

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
func (s *Session) FindAll(db *gorm.DB) (*[]Object, error) {
	entites := []Session{}
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
func (s *Session) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&s).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByID returns an objects with corresponding ID if exists
func (s *Session) FindByID(db *gorm.DB, uid uuid.UUID) error {
	err := db.Model(&s).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates the existing objects
func (s *Session) Update(db *gorm.DB) error {
	if s.ID == uuid.Nil {
		return fmt.Errorf("cannot update non saved session")
	}

	err := s.Validate("update")
	if err != nil {
		return err
	}

	err = db.Model(&s).Updates(Session{
		Name:  s.Name,
		User:  s.User,
		Event: s.Event,
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
func (s *Session) Delete(db *gorm.DB) error {
	err := db.Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}
