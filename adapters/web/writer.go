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

type writer struct {
	engine.Writer
}

func initWriter(f engine.Factory, m *mux.Router) {
	writer := &writer{
		Writer: f.NewWriter(),
	}

	m.HandleFunc(domain.PathWriters, writer.listHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathWriters, writer.createHandler).Methods(http.MethodPost)
	m.HandleFunc(domain.PathWriters+regexInt("id"), writer.findHandler).Methods(http.MethodGet)
	m.HandleFunc(domain.PathWriters+regexInt("id"), writer.updateHandler).Methods(http.MethodPut)
	m.HandleFunc(domain.PathWriters+regexInt("id"), writer.destroyHandler).Methods(http.MethodDelete)

}

func (c *writer) listHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.ListWritersRequest
	if err := form.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.List(&req)
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusOK, res.Writers, paginationHeaders(domain.PathWriters, r.URL.Query(), res.Count))
}

func (c *writer) createHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.CreateWriterRequest
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
	sendJSON(w, http.StatusCreated, res.Writer, noHeader)
}

func (c *writer) findHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Find(&engine.FindWriterRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}

	sendJSON(w, http.StatusOK, res.Writer, noHeader)
}

func (c *writer) updateHandler(w http.ResponseWriter, r *http.Request) {
	var req engine.UpdateWriterRequest
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
	sendJSON(w, http.StatusOK, res.Writer, noHeader)
}

func (c *writer) destroyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	res := c.Destroy(&engine.DestroyWriterRequest{ID: id})
	if res.Error != nil {
		sendErrorJSON(w, res.Error)
		return
	}
	sendJSON(w, http.StatusNoContent, map[string]interface{}{}, noHeader)
}
