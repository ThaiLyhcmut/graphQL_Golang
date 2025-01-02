package models

import "time"

type Category struct {
	ID          uint      `gorm:"primaryKey"`                   // Khóa chính
	Title       string    `gorm:"column:title"`                 // Cột title
	Description string    `gorm:"column:description"`           // Cột description
	Thumbnail   string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status      string    `gorm:"column:status"`                // Cột status
	Postion     uint      `gorm:"column:postion"`               // Cột postion
	Deleted     bool      `gorm:"column:deleted;default:false"` // Cột deleted
	Slug        string    `gorm:"column:slug"`                  // Cột slug
	CreatedAt   time.Time `gorm:"column:createdAt"`             // Cột createdAt
	UpdatedAt   time.Time `gorm:"column:updatedAt"`             // Cột updatedAt
}

func (Category) TableName() string {
	return "categories"
}
