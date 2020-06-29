package model

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Comment represents an user comment in a session
type Comment struct {
	Base
	Message   string `gorm:"size:255;not null" json:"message"`
	User      User   `json:"author"`
	UserID    uuid.UUID
	Session   Session `json:"session"`
	SessionID uuid.UUID
}

// GetID returns the ID
func (c *Comment) GetID() uuid.UUID {
	return c.ID
}

// GetCreatedAt returns the CreatedAt
func (c *Comment) GetCreatedAt() time.Time {
	return c.CreatedAt
}

// SetCreatedAt sets the CreatedAt
func (c *Comment) SetCreatedAt(createdAt time.Time) {
	c.CreatedAt = createdAt
}

// Validate checks structure consistency
func (c *Comment) Validate(action string) error {
	// always check
	if c.Message == "" {
		return fmt.Errorf("required Message")
	}

	if c.User.Name == "" {
		return fmt.Errorf("required Author")
	}

	if c.Session.Name == "" {
		return fmt.Errorf("required Session")
	}

	return nil
}

// Save saves the structure as new object
func (c *Comment) Save(db *gorm.DB) error {
	c.Prepare()
	c.Message = html.EscapeString(strings.TrimSpace(c.Message))

	err := c.Validate("update")
	if err != nil {
		return err
	}

	err = db.Create(&c).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll returns all known objects of this type
func (c *Comment) FindAll(db *gorm.DB) (*[]Object, error) {
	entites := []Comment{}
	err := db.Model(&c).Limit(100).Find(&entites).Error
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
func (c *Comment) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&c).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByID returns an objects with corresponding ID if exists
func (c *Comment) FindByID(db *gorm.DB, uid uuid.UUID) error {
	err := db.Model(&c).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates the existing objects
func (c *Comment) Update(db *gorm.DB) error {
	if c.ID == uuid.Nil {
		return fmt.Errorf("cannot update non saved comment")
	}

	err := c.Validate("update")
	if err != nil {
		return err
	}

	err = db.Model(&c).Updates(Comment{
		Message: c.Message,
		User:    c.User,
		Session: c.Session,
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
func (c *Comment) Delete(db *gorm.DB) error {
	err := db.Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}
