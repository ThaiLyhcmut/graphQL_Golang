package models

import "time"

type Product struct {
	ID              uint      `gorm:"primaryKey"`                   // Khóa chính
	CategoryID      uint      `gorm:"column:categoryID"`            // Cột categoryID
	Title           string    `gorm:"column:title"`                 // Cột title
	Description     string    `gorm:"column:description"`           // Cột description
	Price           float32   `gorm:"column:price"`                 // Cột price
	DiscountPercent int       `gorm:"column:discountPercent"`       // Cột discountPercent
	Stock           int       `gorm:"column:stock"`                 // Cột stock
	Thumbnail       string    `gorm:"column:thumbnail"`             // Cột thumbnail
	Status          string    `gorm:"column:status"`                // Cột status
	Deleted         bool      `gorm:"column:deleted;default:false"` // Cột deleted
	Postion         uint      `gorm:"column:postion"`               // Cột postion
	Slug            string    `gorm:"column:slug"`                  // Cột slug
	Featured        bool      `gorm:"column:featured"`              // Cột featured
	CreatedAt       time.Time `gorm:"column:createdAt"`             // Cột createdAt
	UpdatedAt       time.Time `gorm:"column:updatedAt"`             // Cột updatedAt
	CreatedBy       int       `gorm:"column:createdBy"`             // Cột createdBy
	UpdatedBy       int       `gorm:"column:updatedBy"`             // Cột updatedBy
}
