package engine

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/ttexan1/golang-simple/domain"
)

type (
	// Writer interface defines methods to handle
	// usecases related to writers.
	Writer interface {
		List(*ListWritersRequest) *ListWritersResponse
		Create(*CreateWriterRequest) *CreateWriterResponse
		Find(*FindWriterRequest) *FindWriterResponse
		Update(*UpdateWriterRequest) *UpdateWriterResponse
		Destroy(*DestroyWriterRequest) *DestroyWriterResponse
	}

	writer struct {
		repo WriterRepo
	}
)

func (f *factory) NewWriter() Writer {
	return &writer{
		repo: f.NewWriterRepo(),
	}
}

type (
	// ListWritersRequest is the request
	ListWritersRequest struct {
		Limit  int    `form:"limit"`
		Offset int    `form:"offset"`
		Sort   string `form:"sort"`
	}
	// ListWritersResponse is the response
	ListWritersResponse struct {
		Writers []*domain.Writer
		Count   int
		Error   *domain.Error
	}
)

func (c *writer) List(r *ListWritersRequest) *ListWritersResponse {
	writers, count, err := c.repo.List(r)
	if err != nil {
		return &ListWritersResponse{
			Error: err,
		}
	}
	return &ListWritersResponse{
		Writers: writers,
		Count:   count,
	}
}

type (
	// CreateWriterRequest is the request
	CreateWriterRequest struct {
		DisplayOrder *int   `json:"display_order"`
		Name         string `json:"name"`
	}
	// CreateWriterResponse is the response
	CreateWriterResponse struct {
		Writer *domain.Writer
		Error  *domain.Error
	}
)

func (c *writer) Create(r *CreateWriterRequest) *CreateWriterResponse {
	var params domain.Writer
	if err := copier.Copy(&params, r); err != nil {
		return &CreateWriterResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	writer, err := c.repo.Create(&params)
	return &CreateWriterResponse{
		Writer: writer,
		Error:  err,
	}
}

type (
	// FindWriterRequest is the request
	FindWriterRequest struct {
		ID int `json:"-"`
	}
	// FindWriterResponse is the response
	FindWriterResponse struct {
		Writer *domain.Writer
		Error  *domain.Error
	}
)

func (c *writer) Find(r *FindWriterRequest) *FindWriterResponse {
	writer, err := c.repo.Find(r.ID)
	if err != nil {
		return &FindWriterResponse{
			Error: err,
		}
	}
	return &FindWriterResponse{
		Writer: writer,
	}
}

type (
	// UpdateWriterRequest is the request
	UpdateWriterRequest struct {
		ID           int     `json:"-"`
		DisplayOrder *int    `json:"display_order"`
		Name         *string `json:"name"`
	}
	// UpdateWriterResponse is the response
	UpdateWriterResponse struct {
		Writer *domain.Writer
		Error  *domain.Error
	}
)

func (c *writer) Update(r *UpdateWriterRequest) *UpdateWriterResponse {
	var params domain.Writer
	if err := copier.Copy(&params, r); err != nil {
		return &UpdateWriterResponse{
			Error: domain.NewError(http.StatusInternalServerError, err.Error()),
		}
	}
	writer, err := c.repo.Find(r.ID)
	if err != nil {
		return &UpdateWriterResponse{
			Error: err,
		}
	}
	if err = c.repo.Update(writer, &params); err != nil {
		return &UpdateWriterResponse{
			Error: err,
		}
	}
	return &UpdateWriterResponse{
		Writer: writer,
	}
}

type (
	// DestroyWriterRequest is the request
	DestroyWriterRequest struct {
		ID int `json:"-"`
	}
	// DestroyWriterResponse is the response
	DestroyWriterResponse struct {
		Error *domain.Error
	}
)

func (c *writer) Destroy(r *DestroyWriterRequest) *DestroyWriterResponse {
	err := c.repo.Destroy(r.ID)
	if err != nil {
		return &DestroyWriterResponse{
			Error: err,
		}
	}
	return &DestroyWriterResponse{
		Error: nil,
	}
}
