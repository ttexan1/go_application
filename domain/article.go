package domain

import (
	"time"
)

// Article describes an article
type Article struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CategoryID int `json:"category_id" gorm:"type:integer REFERENCES categories(id); index; not null" valid:"required"`
	WriterID   int `json:"writer_id" gorm:"type:integer REFERENCES writers(id); index; not null" valid:"required"`

	Description *string    `json:"description"`
	ImageURL    *string    `json:"image_url" valid:"url"`
	PublishAt   *time.Time `json:"publish_at" gorm:"index"`
	Status      string     `json:"status" valid:"required,in(draft|public|archived|deleted)"`
	Title       string     `json:"title" valid:"required,runelength(15|48)"`

	Category *Category `json:"category"`
	// Writer          *Writer           `json:"writer"`
}

// article statuses
const (
	ArticleStatusDraft    = "draft"
	ArticleStatusPublic   = "public"
	ArticleStatusArchived = "archived"
	ArticleStatusDeleted  = "deleted"
)
