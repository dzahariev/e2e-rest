package model

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents an user
type User struct {
	Base
	Name          string         `gorm:"size:255;not null;unique" json:"name"`
	Email         string         `gorm:"size:100;not null;unique" json:"email"`
	Password      string         `gorm:"size:100;not null;" json:"password"`
	Subscriptions []Subscription `gorm:"foreignkey:UserID"`
	Sessions      []Session      `gorm:"foreignkey:UserID"`
	Comments      []Comment      `gorm:"foreignkey:UserID"`
}

// GetID returns the ID
func (u *User) GetID() uuid.UUID {
	return u.ID
}

// GetCreatedAt returns the CreatedAt
func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

// SetCreatedAt sets the CreatedAt
func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

// hash hashes the string. Used in password management
func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword checks the known password hash against the hashed string that was provided
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Validate checks structure consistency
func (u *User) Validate(action string) error {
	// always check
	if u.Password == "" {
		return fmt.Errorf("required Password")
	}
	if u.Email == "" {
		return fmt.Errorf("required Email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return fmt.Errorf("invalid Email")
	}

	// specific checks
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return fmt.Errorf("required Name")
		}
	case "login":
		// no specific checks
	default:
		if u.Name == "" {
			return fmt.Errorf("required Name")
		}
	}

	return nil
}

// Save saves the structure as new object
func (u *User) Save(db *gorm.DB) error {
	err := u.Prepare()
	if err != nil {
		return err
	}

	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	err = u.Validate("update")
	if err != nil {
		return err
	}

	hashedPassword, err := hash(u.Password)
	if err != nil {
		return fmt.Errorf("cannot hash the password: %w", err)
	}
	u.Password = string(hashedPassword)

	err = db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll returns all known objects of this type
func (u *User) FindAll(db *gorm.DB) (*[]Object, error) {
	entites := []User{}
	err := db.Model(&u).Limit(100).Find(&entites).Error
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
func (u *User) Count(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&u).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByID returns an objects with corresponding ID if exists
func (u *User) FindByID(db *gorm.DB, uid uuid.UUID) error {
	err := db.Model(&u).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates the existing objects
func (u *User) Update(db *gorm.DB) error {
	if u.ID == uuid.Nil {
		return fmt.Errorf("cannot update non saved user")
	}

	hashedPassword, err := hash(u.Password)
	if err != nil {
		return fmt.Errorf("cannot hash the password: %w", err)
	}
	u.Password = string(hashedPassword)

	err = u.Validate("update")
	if err != nil {
		return err
	}

	err = db.Model(&u).Updates(User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
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
func (u *User) Delete(db *gorm.DB) error {
	err := db.Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}
