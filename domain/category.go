package domain

import "time"

// Category describes a category
type Category struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DisplayOrder int    `json:"display_order" gorm:"default 0"`
	Name         string `json:"name" valid:"required,length(2|20)"`
}
