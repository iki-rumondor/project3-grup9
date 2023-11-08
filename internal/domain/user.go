package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/iki-rumondor/project3-grup9/internal/utils"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"not_null;varchar(120)"`
	Email    string `gorm:"unique;not_null;varchar(120)"`
	Password string `gorm:"not_null;varchar(120)"`
	Role     string `gorm:"not_null;varchar(5)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	if result := tx.First(&User{}, "email = ? AND id != ?", u.Email, u.ID).RowsAffected; result == 1 {
		return errors.New("the email has already been taken")
	}

	validRoles := []string{"admin", "member"}

	if !slices.Contains(validRoles, u.Role) {
		return errors.New("user role is not valid")
	}

	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {

	if result := tx.First(&User{}, "id = ?", u.ID).RowsAffected; result != 1 {
		return fmt.Errorf("user with id %d is not found", u.ID)
	}

	if result := tx.First(&User{}, "email = ? AND id != ?", u.Email, u.ID).RowsAffected; result != 0 {
		return errors.New("the email has already been taken")
	}

	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {

	if result := tx.First(&User{}, "id = ?", u.ID).RowsAffected; result != 1 {
		return fmt.Errorf("user with id %d is not found", u.ID)
	}

	return nil
}
