package testseed

import (
	"time"

	"github.com/ttexan1/golang-simple/domain"
)

// Articles are the seed data
var Articles = []*domain.Article{
	&domain.Article{
		ID: 1,
		// CategoryID:  Categories[0].ID,
		// WriterID:    Writers[0].ID,
		Description: domain.PtrString("friend"),
		ImageURL:    domain.PtrString("http://image.org"),
		PublishAt:   domain.PtrTime(time.Now().Add(time.Hour * time.Duration(0))),
		Status:      domain.ArticleStatusPublic,
		Title:       "This is one of the greatest article",
	},
	&domain.Article{
		ID: 2,
		// CategoryID:  Categories[0].ID,
		// WriterID:    Writers[0].ID,
		Description: domain.PtrString("friend"),
		ImageURL:    domain.PtrString("http://image.org"),
		PublishAt:   domain.PtrTime(time.Now().Add(time.Hour * time.Duration(1))),
		Status:      domain.ArticleStatusPublic,
		Title:       "This is great article",
	},
	&domain.Article{
		ID: 3,
		// CategoryID:  Categories[0].ID,
		// WriterID:    Writers[0].ID,
		Description: domain.PtrString("friend"),
		ImageURL:    domain.PtrString("http://image.org"),
		PublishAt:   domain.PtrTime(time.Now().Add(time.Hour * time.Duration(2))),
		Status:      domain.ArticleStatusPublic,
		Title:       "This is great article",
	},
}
