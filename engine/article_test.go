package engine

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/ttexan1/golang-simple/domain"
)

type ArticleSuite struct {
	suite.Suite
	factory Factory
	article Article
}

type testArticleRepo struct{}

func (r *testArticleRepo) List(params *ListArticlesRequest) ([]*domain.Article, int, *domain.Error) {
	if params.Limit == 1 {
		return []*domain.Article{
			&domain.Article{},
		}, 1, nil
	}
	return nil, 0, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testArticleRepo) Create(c *domain.Article) (*domain.Article, *domain.Error) {
	if c.Title == "test" {
		return c, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testArticleRepo) Find(id int) (*domain.Article, *domain.Error) {
	if id == 1 {
		return &domain.Article{}, nil
	}
	return nil, &domain.Error{
		Code: http.StatusBadRequest,
	}
}
func (r *testArticleRepo) Update(article, params *domain.Article) *domain.Error {
	if article.ID == 1 {
		return nil
	}
	return &domain.Error{Code: http.StatusBadRequest}
}
func (r *testArticleRepo) Destroy(id int) *domain.Error {
	if id == 1 {
		return nil
	}
	return &domain.Error{Code: http.StatusBadRequest}
}

func TestArticleSuite(t *testing.T) {
	suite.Run(t, &ArticleSuite{
		article: &article{
			repo: &testArticleRepo{},
		},
		factory: &factory{
			StorageFactory: &testStorage{},
		},
	})
}

func (s *ArticleSuite) TestNewArticle() {
	s.NotNil(s.factory.NewArticle())
}

func (s *ArticleSuite) TestListError() {
	resp := s.article.List(&ListArticlesRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *ArticleSuite) TestList() {
	resp := s.article.List(&ListArticlesRequest{
		Limit: 1,
	})
	s.Equal(1, resp.Count)
}

func (s *ArticleSuite) TestCreateError() {
	resp := s.article.Create(&CreateArticleRequest{
		WriterClaims: &domain.JWTClaims{
			WriterID: 1,
		},
	})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *ArticleSuite) TestCreate() {
	resp := s.article.Create(&CreateArticleRequest{
		Title: "test",
		WriterClaims: &domain.JWTClaims{
			WriterID: 1,
		},
	})
	s.Nil(resp.Error)
	s.Equal(1, resp.Article.WriterID)
}

func (s *ArticleSuite) TestFindError() {
	resp := s.article.Find(&FindArticleRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *ArticleSuite) TestFind() {
	resp := s.article.Find(&FindArticleRequest{
		ID: 1,
	})
	s.NotNil(resp.Article)
}

func (s *ArticleSuite) TestUpdateError() {
	resp := s.article.Update(&UpdateArticleRequest{ID: 1})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *ArticleSuite) TestDestroyError() {
	resp := s.article.Destroy(&DestroyArticleRequest{})
	s.Equal(http.StatusBadRequest, resp.Error.Code)
}

func (s *ArticleSuite) TestDestroy() {
	resp := s.article.Destroy(&DestroyArticleRequest{
		ID: 1,
	})
	s.Nil(resp.Error)
}
