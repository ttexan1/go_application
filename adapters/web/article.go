package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

type article struct {
	engine.Article
}

func initArticle(f engine.Factory, m *mux.Router) {
	article := &article{
		Article: f.NewArticle(),
	}

	m.HandleFunc(domain.PathArticles, article.listHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathArticles, article.createHandler).Methods(http.MethodPost)
	m.HandleFunc(domain.PathArticles+regexInt("id"), article.findHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathArticles+regexInt("id"), article.updateHandler).Methods(http.MethodPut)
	m.HandleFunc(domain.PathArticles+regexInt("id"), article.destroyHandler).Methods(http.MethodDelete)

}

func (c *article) listHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.ListArticlesRequest
	if err := form.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.List(&req)
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusOK, res.Articles, paginationHeaders(domain.PathArticles, r.URL.Query(), res.Count))
}

func (c *article) createHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.CreateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Create(&req)
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusCreated, res.Article, noHeader)
}

func (c *article) findHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Find(&engine.FindArticleRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}

	sendJSON(w, http.StatusOK, res.Article, noHeader)
}

func (c *article) updateHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.UpdateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	req.ID, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Update(&req)
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusOK, res.Article, noHeader)
}

func (c *article) destroyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Destroy(&engine.DestroyArticleRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusNoContent, map[string]interface{}{}, noHeader)
}
