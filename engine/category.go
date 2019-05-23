package engine

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/ttexan1/golang-simple/domain"
)

type (
	// Category interface defines methods to handle
	// usecases related to categories.
	Category interface {
		List(*ListCategoriesRequest) *ListCategoriesResponse
		Create(*CreateCategoryRequest) *CreateCategoryResponse
		Find(*FindCategoryRequest) *FindCategoryResponse
		Update(*UpdateCategoryRequest) *UpdateCategoryResponse
		Destroy(*DestroyCategoryRequest) *DestroyCategoryResponse
	}

	category struct {
		repo CategoryRepo
	}
)

func (f *factory) NewCategory() Category {
	return &category{
		repo: f.NewCategoryRepo(),
	}
}

type (
	// ListCategoriesRequest is the request
	ListCategoriesRequest struct {
		Limit  int    `form:"limit"`
		Offset int    `form:"offset"`
		Sort   string `form:"sort"`
	}
	// ListCategoriesResponse is the response
	ListCategoriesResponse struct {
		Categories []*domain.Category
		Count      int
		Error      *domain.Error
	}
)

func (c *category) List(r *ListCategoriesRequest) *ListCategoriesResponse {
	categories, count, err := c.repo.List(r)
	if err != nil {
		return &ListCategoriesResponse{
			Error: err,
		}
	}
	return &ListCategoriesResponse{
		Categories: categories,
		Count:      count,
	}
}

type (
	// CreateCategoryRequest is the request
	CreateCategoryRequest struct {
		DisplayOrder *int   `json:"display_order"`
		Name         string `json:"name"`
	}
	// CreateCategoryResponse is the response
	CreateCategoryResponse struct {
		Category *domain.Category
		Error    *domain.Error
	}
)

func (c *category) Create(r *CreateCategoryRequest) *CreateCategoryResponse {
	var params domain.Category
	if err := copier.Copy(&params, r); err != nil {
		return &CreateCategoryResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	category, err := c.repo.Create(&params)
	return &CreateCategoryResponse{
		Category: category,
		Error:    err,
	}
}

type (
	// FindCategoryRequest is the request
	FindCategoryRequest struct {
		ID int `json:"-"`
	}
	// FindCategoryResponse is the response
	FindCategoryResponse struct {
		Category *domain.Category
		Error    *domain.Error
	}
)

func (c *category) Find(r *FindCategoryRequest) *FindCategoryResponse {
	category, err := c.repo.Find(r.ID)
	if err != nil {
		return &FindCategoryResponse{
			Error: err,
		}
	}
	return &FindCategoryResponse{
		Category: category,
	}
}

type (
	// UpdateCategoryRequest is the request
	UpdateCategoryRequest struct {
		ID           int     `json:"-"`
		DisplayOrder *int    `json:"display_order"`
		Name         *string `json:"name"`
	}
	// UpdateCategoryResponse is the response
	UpdateCategoryResponse struct {
		Category *domain.Category
		Error    *domain.Error
	}
)

func (c *category) Update(r *UpdateCategoryRequest) *UpdateCategoryResponse {
	var params domain.Category
	if err := copier.Copy(&params, r); err != nil {
		return &UpdateCategoryResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	category, err := c.repo.Find(r.ID)
	if err != nil {
		return &UpdateCategoryResponse{
			Error: err,
		}
	}
	err = c.repo.Update(category, &params)
	return &UpdateCategoryResponse{
		Category: category,
		Error:    err,
	}
}

type (
	// DestroyCategoryRequest is the request
	DestroyCategoryRequest struct {
		ID int `json:"-"`
	}
	// DestroyCategoryResponse is the response
	DestroyCategoryResponse struct {
		Error *domain.Error
	}
)

func (c *category) Destroy(r *DestroyCategoryRequest) *DestroyCategoryResponse {
	err := c.repo.Destroy(r.ID)
	if err != nil {
		return &DestroyCategoryResponse{
			Error: err,
		}
	}
	return &DestroyCategoryResponse{
		Error: nil,
	}
}
