package psql

import (
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

const tblCategories = "categories"

type categoryRepo struct {
	db *string
}

func newCategoryRepo(db *string) *categoryRepo {
	return &categoryRepo{db}
}

func (r *categoryRepo) Destroy(id int) (err *domain.Error) {
	return
}

func (r *categoryRepo) List(params *engine.ListCategoriesRequest) (categories []*domain.Category, count int, err *domain.Error) {
	categories = []*domain.Category{
		&domain.Category{
			ID:           1,
			Name:         "Category1",
			DisplayOrder: 1,
		},
	}
	return
}

func (r *categoryRepo) Find(id int) (category *domain.Category, err *domain.Error) {
	return
}

func (r *categoryRepo) Create(params *domain.Category) (*domain.Category, *domain.Error) {
	return r.Find(params.ID)
}

func (r *categoryRepo) Update(category, params *domain.Category) *domain.Error {
	return &domain.Error{}
}
