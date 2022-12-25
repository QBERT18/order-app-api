package entity

import "time"

type Product struct {
	ID          uint            `gorm:"primary_key" json:"id"`
	Name        string          `gorm:"type:varchar(255)" json:"name"`
	Description string          `gorm:"type:text" json:"description"`
	Price       float64         `json:"price"`
	CategoryID  uint            `gorm:"index" json:"category_id"`
	Category    ProductCategory `gorm:"foreignkey:CategoryID;association_foreignkey:ID"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at"`
}
