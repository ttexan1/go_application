package sql

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

const tblArticles = "articles"

type articleRepo struct {
	db *gorm.DB
}

func newArticleRepo(db *gorm.DB) *articleRepo {
	return &articleRepo{db}
}

func (r *articleRepo) Destroy(id int) (err *domain.Error) {
	gormErr := r.db.Table(tblArticles).
		Delete(domain.Article{}, "id = ?", id).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *articleRepo) List(params *engine.ListArticlesRequest) (articles []*domain.Article, count int, err *domain.Error) {
	db := r.db.Table(tblArticles)
	if params.CategoryID != 0 {
		db = db.Where("articles.category_id = ?", params.CategoryID)
	}
	if params.PublishAtGteq != nil {
		db = db.Where("articles.publish_at >= ?", *params.PublishAtGteq)
	}
	if params.PublishAtLteq != nil {
		db = db.Where("articles.publish_at <= ?", *params.PublishAtLteq)
	}
	if params.Title != "" {
		for _, w := range strings.Fields(params.Title) {
			db = db.Where("UPPER(articles.title) LIKE ?", "%"+strings.ToUpper(w)+"%")
		}
	}
	if len(params.Statuses) != 0 {
		db = db.Where("articles.status in (?)", params.Statuses)
	}
	if params.WriterID != 0 {
		db = db.Where("articles.writer_id = ?", params.WriterID)
	}
	db = db.Count(&count)
	db = sort(db, params.Sort, []string{"id", "display_order"}, "id DESC")
	db = limit(db, params.Limit, 20)
	db = offset(db, params.Offset)
	gormErr := db.Find(&articles).Error
	err = newErrorByGormError(gormErr)
	return
}

func (r *articleRepo) Find(id int) (article *domain.Article, err *domain.Error) {
	article = &domain.Article{}
	gormErr := r.db.Table(tblArticles).
		Where(domain.Article{ID: id}).
		First(&article).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *articleRepo) Create(params *domain.Article) (*domain.Article, *domain.Error) {
	if err := newErrorByGormError(
		r.db.Table(tblArticles).
			Create(&params).Error); err != nil {
		return nil, err
	}
	return r.Find(params.ID)
}

func (r *articleRepo) Update(article, params *domain.Article) *domain.Error {
	return newErrorByGormError(
		r.db.Model(article).
			Updates(params).Error)
}
