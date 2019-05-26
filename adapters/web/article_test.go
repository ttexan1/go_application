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

type ArticleSuite struct {
	webSuite
}

type testArticleEngine struct{}

func (e *testArticleEngine) Find(r *engine.FindArticleRequest) *engine.FindArticleResponse {
	if r.ID == 1 {
		return &engine.FindArticleResponse{
			Article: &domain.Article{ID: 1},
		}
	}
	return &engine.FindArticleResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testArticleEngine) List(r *engine.ListArticlesRequest) *engine.ListArticlesResponse {
	return &engine.ListArticlesResponse{
		Error: &domain.Error{Code: http.StatusNotFound},
	}
}

func (e *testArticleEngine) Create(r *engine.CreateArticleRequest) *engine.CreateArticleResponse {
	return &engine.CreateArticleResponse{
		Article: &domain.Article{
			ID:     1,
			Title:  "article",
			Status: domain.ArticleStatusDraft,
		},
		Error: nil,
	}
}

func (e *testArticleEngine) Update(r *engine.UpdateArticleRequest) *engine.UpdateArticleResponse {
	if r.ID == 1 {
		return &engine.UpdateArticleResponse{
			Article: &domain.Article{
				ID:     1,
				Title:  "update article",
				Status: domain.ArticleStatusDraft,
			},
			Error: nil,
		}
	}
	return &engine.UpdateArticleResponse{
		Article: &domain.Article{},
		Error:   domain.NewError(http.StatusNotFound, gorm.ErrRecordNotFound.Error()),
	}
}

func (e *testArticleEngine) Destroy(r *engine.DestroyArticleRequest) *engine.DestroyArticleResponse {
	return &engine.DestroyArticleResponse{Error: nil}
}

func TestArticleSuite(t *testing.T) {
	m := mux.NewRouter()
	server := httptest.NewServer(m)
	defer server.Close()
	initArticle(&testFactory{}, m)
	suite.Run(t, &ArticleSuite{
		webSuite{
			server: server,
		},
	})
}

func (s *ArticleSuite) TestListHandlerError() {
	s.doRequest(http.MethodGet, domain.PathArticles, map[string]string{}, "").
		checkStatus(http.StatusNotFound)
}

func (s *ArticleSuite) TestFindHandler() {
	s.doRequest(http.MethodGet, domain.PathArticles+"/0", map[string]string{}, "").
		checkStatus(http.StatusNotFound)
	var resp engine.FindArticleResponse
	s.get(domain.PathArticles+"/1", url.Values{}, &resp.Article, "")
}

func (s *ArticleSuite) TestCreateHandler() {
	s.doRequest(http.MethodPost, domain.PathArticles, map[string]string{}, "").
		checkStatus(http.StatusUnauthorized)
}
