package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
	"github.com/ttexan1/golang-simple/providers/sql/testdata"
)

type ArticleSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *articleRepo
}

func (s *ArticleSuite) SetupTest() {
	cleaner.Acquire(tableNames...)
	s.repo = newArticleRepo(s.db)
	for _, c := range testdata.Writers {
		c.SetPassword("password")
		_, err := newWriterRepo(s.db).Create(c)
		s.Require().Nil(err)
	}
	for _, c := range testdata.Categories {
		_, err := newCategoryRepo(s.db).Create(c)
		s.Require().Nil(err)
	}
	for _, c := range testdata.Articles {
		_, err := s.repo.Create(c)
		s.Require().Nil(err)
	}
}

func (s *ArticleSuite) TearDownTest() {
	cleaner.Clean(tableNames...)
}

func (s *ArticleSuite) TestDestroy() {
	s.Nil(s.repo.Destroy(1))
	_, err := s.repo.Find(1)
	s.NotNil(err)
}

func (s *ArticleSuite) TestListSort() {
	list, count, err := s.repo.List(&engine.ListArticlesRequest{
		Limit: 1,
		Sort:  "id asc",
	})
	s.Nil(err)
	s.Equal(1, len(list))
	s.Equal(len(testdata.Articles), count)
}

func (s *ArticleSuite) TestListWriterID() {
	list, _, err := s.repo.List(&engine.ListArticlesRequest{
		WriterID: 1,
		Sort:     "id desc",
	})
	s.Nil(err)
	s.Equal(2, len(list))
	s.Equal(testdata.Articles[1].Title, list[0].Title)
}

func (s *ArticleSuite) TestListOffset() {
	list, count, err := s.repo.List(&engine.ListArticlesRequest{
		Offset: 1,
	})
	s.Nil(err)
	s.Equal(2, len(list))
	s.Equal(len(testdata.Articles), count)
}

func (s *ArticleSuite) TestFind() {
	a, err := s.repo.Find(1)
	s.Nil(err)
	s.Equal(testdata.Articles[0].Title, a.Title)
}

func (s *ArticleSuite) TestUpdate() {
	err := s.repo.Update(testdata.Articles[0], &domain.Article{
		ID:    1,
		Title: "this is updated title",
	})
	s.Nil(err)
	s.Equal("this is updated title", testdata.Articles[0].Title)
}
