package engine

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/ttexan1/golang-simple/domain"
)

type (
	// Article interface defines methods to handle
	// usecases related to articles.
	Article interface {
		List(*ListArticlesRequest) *ListArticlesResponse
		Create(*CreateArticleRequest) *CreateArticleResponse
		Find(*FindArticleRequest) *FindArticleResponse
		Update(*UpdateArticleRequest) *UpdateArticleResponse
		Destroy(*DestroyArticleRequest) *DestroyArticleResponse
	}

	article struct {
		repo ArticleRepo
	}
)

func (f *factory) NewArticle() Article {
	return &article{
		repo: f.NewArticleRepo(),
	}
}

type (
	// ListArticlesRequest is the request
	ListArticlesRequest struct {
		Limit  int    `form:"limit"`
		Offset int    `form:"offset"`
		Sort   string `form:"sort"`
	}
	// ListArticlesResponse is the response
	ListArticlesResponse struct {
		Articles []*domain.Article
		Count    int
		Error    *domain.Error
	}
)

func (c *article) List(r *ListArticlesRequest) *ListArticlesResponse {
	articles, count, err := c.repo.List(r)
	if err != nil {
		return &ListArticlesResponse{
			Error: err,
		}
	}
	return &ListArticlesResponse{
		Articles: articles,
		Count:    count,
	}
}

type (
	// CreateArticleRequest is the request
	CreateArticleRequest struct {
		DisplayOrder *int   `json:"display_order"`
		Name         string `json:"name"`
	}
	// CreateArticleResponse is the response
	CreateArticleResponse struct {
		Article *domain.Article
		Error   *domain.Error
	}
)

func (c *article) Create(r *CreateArticleRequest) *CreateArticleResponse {
	var params domain.Article
	if err := copier.Copy(&params, r); err != nil {
		return &CreateArticleResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	article, err := c.repo.Create(&params)
	return &CreateArticleResponse{
		Article: article,
		Error:   err,
	}
}

type (
	// FindArticleRequest is the request
	FindArticleRequest struct {
		ID int `json:"-"`
	}
	// FindArticleResponse is the response
	FindArticleResponse struct {
		Article *domain.Article
		Error   *domain.Error
	}
)

func (c *article) Find(r *FindArticleRequest) *FindArticleResponse {
	article, err := c.repo.Find(r.ID)
	if err != nil {
		return &FindArticleResponse{
			Error: err,
		}
	}
	return &FindArticleResponse{
		Article: article,
	}
}

type (
	// UpdateArticleRequest is the request
	UpdateArticleRequest struct {
		ID           int     `json:"-"`
		DisplayOrder *int    `json:"display_order"`
		Name         *string `json:"name"`
	}
	// UpdateArticleResponse is the response
	UpdateArticleResponse struct {
		Article *domain.Article
		Error   *domain.Error
	}
)

func (c *article) Update(r *UpdateArticleRequest) *UpdateArticleResponse {
	var params domain.Article
	if err := copier.Copy(&params, r); err != nil {
		return &UpdateArticleResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	article, err := c.repo.Find(r.ID)
	if err != nil {
		return &UpdateArticleResponse{
			Error: err,
		}
	}
	if err = c.repo.Update(article, &params); err != nil {
		return &UpdateArticleResponse{
			Error: err,
		}
	}
	return &UpdateArticleResponse{
		Article: article,
	}
}

type (
	// DestroyArticleRequest is the request
	DestroyArticleRequest struct {
		ID int `json:"-"`
	}
	// DestroyArticleResponse is the response
	DestroyArticleResponse struct {
		Error *domain.Error
	}
)

func (c *article) Destroy(r *DestroyArticleRequest) *DestroyArticleResponse {
	err := c.repo.Destroy(r.ID)
	if err != nil {
		return &DestroyArticleResponse{
			Error: err,
		}
	}
	return &DestroyArticleResponse{
		Error: nil,
	}
}
