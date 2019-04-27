package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

const tblCategories = "categories"

type categoryRepo struct {
	db *gorm.DB
}

func newCategoryRepo(db *gorm.DB) *categoryRepo {
	return &categoryRepo{db}
}

func (r *categoryRepo) Destroy(id int) (err *domain.Error) {
	gormErr := r.db.Table(tblCategories).
		Delete(domain.Category{}, "id = ?", id).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *categoryRepo) List(params *engine.ListCategoriesRequest) (categories []*domain.Category, count int, err *domain.Error) {
	db := r.db.Table(tblCategories)
	db = db.Count(&count)
	db = sort(db, params.Sort, []string{"id", "display_order"}, "id DESC")
	db = limit(db, params.Limit, 20)
	db = offset(db, params.Offset)
	gormErr := db.Find(&categories).Error
	err = newErrorByGormError(gormErr)
	return
}

func (r *categoryRepo) Find(id int) (category *domain.Category, err *domain.Error) {
	category = &domain.Category{}
	gormErr := r.db.Table(tblCategories).
		Where(domain.Category{ID: id}).
		First(&category).Error

	err = newErrorByGormError(gormErr)
	return
}

func (r *categoryRepo) Create(params *domain.Category) (*domain.Category, *domain.Error) {
	if err := newErrorByGormError(
		r.db.Table(tblCategories).
			Create(&params).Error); err != nil {
		return nil, err
	}
	return r.Find(params.ID)
}

func (r *categoryRepo) Update(category, params *domain.Category) *domain.Error {
	return newErrorByGormError(
		r.db.Model(category).
			Updates(params).Error)
}
