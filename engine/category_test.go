package engine

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
)

type CategorySuite struct {
	suite.Suite
	factory  Factory
	category Category
}

type testCategoryRepo struct{}

func (r *testCategoryRepo) List(params *ListCategoriesRequest) ([]*domain.Category, int, *domain.Error) {
	if params.Limit == 1 {
		return []*domain.Category{
			&domain.Category{},
		}, 1, nil
	}
	return nil, 0, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testCategoryRepo) Create(c *domain.Category) (*domain.Category, *domain.Error) {
	if c.Name == "test" {
		return c, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testCategoryRepo) Find(id int) (*domain.Category, *domain.Error) {
	if id == 1 {
		return &domain.Category{}, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testCategoryRepo) Update(category, params *domain.Category) *domain.Error {
	if category.ID == 1 {
		return nil
	}
	return &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testCategoryRepo) Destroy(id int) *domain.Error {
	if id == 1 {
		return nil
	}
	return &domain.Error{
		Code: http.StatusBadRequest,
	}
}

func TestCategorySuite(t *testing.T) {
	suite.Run(t, &CategorySuite{
		category: &category{
			repo: &testCategoryRepo{},
		},
		factory: &factory{
			StorageFactory: &testStorage{},
		},
	})
}

func (s *CategorySuite) TestNewCategory() {
	s.NotNil(s.factory.NewCategory())
}

func (s *CategorySuite) TestListError() {
	resp := s.category.List(&ListCategoriesRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *CategorySuite) TestList() {
	resp := s.category.List(&ListCategoriesRequest{
		Limit: 1,
	})
	s.Equal(1, resp.Count)
}

func (s *CategorySuite) TestCreateError() {
	resp := s.category.Create(&CreateCategoryRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *CategorySuite) TestCreate() {
	resp := s.category.Create(&CreateCategoryRequest{
		Name: "test",
	})
	s.Nil(resp.Error)
}

func (s *CategorySuite) TestFindError() {
	resp := s.category.Find(&FindCategoryRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *CategorySuite) TestFind() {
	resp := s.category.Find(&FindCategoryRequest{
		ID: 1,
	})
	s.NotNil(resp.Category)
}

func (s *CategorySuite) TestUpdateError() {
	resp := s.category.Update(&UpdateCategoryRequest{ID: 1})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *CategorySuite) TestDestroyError() {
	resp := s.category.Destroy(&DestroyCategoryRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *CategorySuite) TestDestroy() {
	resp := s.category.Destroy(&DestroyCategoryRequest{
		ID: 1,
	})
	s.Nil(resp.Error)
}
