package model

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Event represents an event
type Event struct {
	Base
	Name     string    `gorm:"size:255;not null;unique" json:"name"`
	Year     string    `gorm:"size:4;not null" json:"year"`
	Sessions []Session `gorm:"foreignkey:EventID"`
}

// GetID returns the ID
func (e *Event) GetID() uuid.UUID {
	return e.ID
}

// GetCreatedAt returns the CreatedAt
func (e *Event) GetCreatedAt() time.Time {
	return e.CreatedAt
}

// SetCreatedAt sets the CreatedAt
func (e *Event) SetCreatedAt(createdAt time.Time) {
	e.CreatedAt = createdAt
}

// Validate checks structure consistency
func (e *Event) Validate(action string) error {
	// always check
	if e.Name == "" {
		return fmt.Errorf("required Name")
	}
	if e.Name == "" {
		return fmt.Errorf("required Year")
	}

	return nil
}

// Save saves the structure as new object
func (e *Event) Save(db *gorm.DB) error {
	err := e.Prepare()
	if err != nil {
		return err
	}

	e.Name = html.EscapeString(strings.TrimSpace(e.Name))
	e.Year = html.EscapeString(strings.TrimSpace(e.Year))

	err = e.Validate("update")
	if err != nil {
		return err
	}

	err = db.Create(&e).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns all known objects of this type
func (e *Event) FindAll(db *gorm.DB) (*[]Object, error) {
	entites := []Event{}
	err := db.Model(&e).Limit(100).Find(&entites).Error
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
func (e *Event) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&e).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByID returns an objects with corresponding ID if exists
func (e *Event) FindByID(db *gorm.DB, uid uuid.UUID) error {
	err := db.Model(&e).Where("id = ?", uid).Take(&e).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates the existing objects
func (e *Event) Update(db *gorm.DB) error {
	if e.ID == uuid.Nil {
		return fmt.Errorf("cannot update non saved event")
	}

	err := e.Validate("update")
	if err != nil {
		return err
	}

	err = db.Model(&e).Updates(Event{
		Name: e.Name,
		Year: e.Year,
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
func (e *Event) Delete(db *gorm.DB) error {
	err := db.Delete(&e).Error
	if err != nil {
		return err
	}
	return nil
}
