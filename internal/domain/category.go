package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not_null;varchar(120)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []Task
}

func (c *Category) BeforeUpdate(tx *gorm.DB) error {

	if result := tx.First(&Category{}, "id = ?", c.ID).RowsAffected; result != 1 {
		return fmt.Errorf("category with id %d is not found", c.ID)
	}

	return nil
}

func (c *Category) BeforeDelete(tx *gorm.DB) error {

	if result := tx.First(&Category{}, "id = ?", c.ID).RowsAffected; result != 1 {
		return fmt.Errorf("category with id %d is not found", c.ID)
	}

	return nil
}
