package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
	"github.com/ttexan1/golang-simple/providers/sql/testdata"
)

type CategorySuite struct {
	suite.Suite
	db   *gorm.DB
	repo *categoryRepo
}

func (s *CategorySuite) SetupTest() {
	cleaner.Acquire(tableNames...)
	s.repo = newCategoryRepo(s.db)
	for _, c := range testdata.Categories {
		_, err := s.repo.Create(c)

		s.Require().Nil(err)
	}
}

func (s *CategorySuite) TearDownTest() {
	cleaner.Clean(tableNames...)
}

func (s *CategorySuite) TestDestroy() {
	s.Nil(s.repo.Destroy(1))
	_, err := s.repo.Find(1)
	s.NotNil(err)
}

func (s *CategorySuite) TestList() {
	list, count, err := s.repo.List(&engine.ListCategoriesRequest{})
	s.Nil(err)
	s.Equal(len(testdata.Categories), len(list))
	s.Equal(len(testdata.Categories), count)
}

func (s *CategorySuite) TestListLimit() {
	list, count, err := s.repo.List(&engine.ListCategoriesRequest{
		Limit: 1,
	})
	s.Nil(err)
	s.Equal(1, len(list))
	s.Equal(len(testdata.Categories), count)
}

func (s *CategorySuite) TestListSort() {
	list, count, err := s.repo.List(&engine.ListCategoriesRequest{
		Sort: "id",
	})
	s.Nil(err)
	s.Equal(len(testdata.Categories), len(list))
	s.Equal(len(testdata.Categories), count)
}

func (s *CategorySuite) TestListOffset() {
	list, count, err := s.repo.List(&engine.ListCategoriesRequest{
		Offset: 1,
	})
	s.Nil(err)
	s.Equal(2, len(list))
	s.Equal(len(testdata.Categories), count)
}

func (s *CategorySuite) TestFind() {
	_, err := s.repo.Find(1)
	s.Nil(err)
}

func (s *CategorySuite) TestUpdate() {
	err := s.repo.Update(testdata.Categories[0], &domain.Category{
		ID:   1,
		Name: "updated",
	})
	s.Nil(err)
	s.Equal("updated", testdata.Categories[0].Name)
}
