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

type category struct {
	engine.Category
}

func initCategory(f engine.Factory, m *mux.Router) {
	category := &category{
		Category: f.NewCategory(),
	}

	m.HandleFunc(domain.PathCategories, category.listHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathCategories, category.createHandler).Methods(http.MethodPost)
	m.HandleFunc(domain.PathCategories+regexInt("id"), category.findHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathCategories+regexInt("id"), category.updateHandler).Methods(http.MethodPut)
	m.HandleFunc(domain.PathCategories+regexInt("id"), category.destroyHandler).Methods(http.MethodDelete)

}

func (c *category) listHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.ListCategoriesRequest
	if err := form.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.List(&req)
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusOK, res.Categories, paginationHeaders(domain.PathCategories, r.URL.Query(), res.Count))
}

func (c *category) createHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.CreateCategoryRequest
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
	sendJSON(w, http.StatusCreated, res.Category, noHeader)
}

func (c *category) findHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Find(&engine.FindCategoryRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}

	sendJSON(w, http.StatusOK, res.Category, noHeader)
}

func (c *category) updateHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.UpdateCategoryRequest
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
	sendJSON(w, http.StatusOK, res.Category, noHeader)
}

func (c *category) destroyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Destroy(&engine.DestroyCategoryRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusNoContent, map[string]interface{}{}, noHeader)
}
