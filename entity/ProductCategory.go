package entity

import "time"

type ProductCategory struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `gorm:"type:varchar(255)" json:"name"`
	Image       string     `gorm:"type:varchar(255)" json:"image"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
