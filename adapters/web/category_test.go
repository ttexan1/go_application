package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

type CategorySuite struct {
	webSuite
}

type testCategoryEngine struct{}

func (e *testCategoryEngine) Find(r *engine.FindCategoryRequest) *engine.FindCategoryResponse {
	if r.ID == 1 {
		return &engine.FindCategoryResponse{
			Category: &domain.Category{ID: 1},
		}
	}
	return &engine.FindCategoryResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testCategoryEngine) List(r *engine.ListCategoriesRequest) *engine.ListCategoriesResponse {
	return &engine.ListCategoriesResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testCategoryEngine) Create(r *engine.CreateCategoryRequest) *engine.CreateCategoryResponse {
	return &engine.CreateCategoryResponse{
		Category: &domain.Category{
			ID:           1,
			Name:         "category",
			DisplayOrder: domain.PtrInt(1),
		},
		Error: nil,
	}
}

func (e *testCategoryEngine) Update(r *engine.UpdateCategoryRequest) *engine.UpdateCategoryResponse {
	if r.ID == 1 {
		return &engine.UpdateCategoryResponse{
			Category: &domain.Category{
				ID:           1,
				Name:         "update category",
				DisplayOrder: domain.PtrInt(1),
			},
			Error: nil,
		}
	}
	return &engine.UpdateCategoryResponse{
		Category: &domain.Category{},
		Error:    domain.NewError(http.StatusNotFound, gorm.ErrRecordNotFound.Error()),
	}
}

func (e *testCategoryEngine) Destroy(r *engine.DestroyCategoryRequest) *engine.DestroyCategoryResponse {
	return &engine.DestroyCategoryResponse{Error: nil}
}

func TestCategorySuite(t *testing.T) {
	m := mux.NewRouter()
	server := httptest.NewServer(m)
	defer server.Close()
	initCategory(&testFactory{}, m)
	suite.Run(t, &CategorySuite{
		webSuite{
			server: server,
		},
	})
}

func (s *CategorySuite) TestListHandlerError() {
	s.doRequest(http.MethodGet, domain.PathCategories, map[string]string{}, "").
		checkStatus(http.StatusNotFound)
}

func (s *CategorySuite) TestFindHandler() {
	s.doRequest(http.MethodGet, domain.PathCategories+"/0", map[string]string{}, "").
		checkStatus(http.StatusNotFound)
	var resp engine.FindCategoryResponse
	s.get(domain.PathCategories+"/1", url.Values{}, &resp.Category, "")
}

func (s *CategorySuite) TestCreateHandler() {
	s.doRequest(http.MethodPost, domain.PathCategories, map[string]string{}, "").
		checkStatus(http.StatusCreated)
}
