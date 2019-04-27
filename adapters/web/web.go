package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
)

var store *sessions.CookieStore
var baseURL string
var noHeader = map[string]string{}

// NewAdapter returns an adapter
func NewAdapter(f engine.Factory, config *domain.Config) http.Handler {
	store = sessions.NewCookieStore([]byte(config.SessionSecret))
	baseURL = config.BaseURL()
	m := mux.NewRouter()
	initCategory(f, m)
	if config.IsProd() {
		return m
	}
	initDocs(m)
	if config.IsSandbox() {
		return m
	}
	return m
	// handler := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut,
	// 	},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: true,
	// 	ExposedHeaders: []string{
	// 		"X-Total-Count",
	// 	},
	// }).Handler(m)
	// fmt.Println("")
	// return handler
}

func paginationHeaders(path string, query url.Values, count int) map[string]string {
	offset, _ := strconv.Atoi(query.Get("offset"))
	if offset == 0 {
		offset = 0
	}
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit == 0 {
		limit = 20
	}
	nextOffset := offset + limit
	prevOffset := offset - limit
	if prevOffset < 0 {
		prevOffset = 0
	}

	query.Set("offset", fmt.Sprintf("%d", nextOffset))
	nextPage := baseURL + path + "?" + query.Encode()
	if nextOffset >= count {
		nextPage = ""
	}

	query.Set("offset", fmt.Sprintf("%d", prevOffset))
	prevPage := baseURL + path + "?" + query.Encode()
	if offset == 0 {
		prevPage = ""
	}

	headers := map[string]string{
		"X-Total-Count": fmt.Sprintf("%d", count),
		"X-Next-Page":   nextPage,
		"X-Prev-Page":   prevPage,
	}
	return headers
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}, headers map[string]string) {
	bytes, err := json.Marshal(data)
	if err != nil {
		sendErrorJSON(w, domain.NewError(http.StatusInternalServerError, err.Error()))
		return
	}
	log.Printf("[INFO]: %s\n", string(bytes))
	w.Header().Set("Content-Type", "application/json")
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func sendErrorJSON(w http.ResponseWriter, err *domain.Error) {
	log.Printf("[ERROR]: %#v\n", err)
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(struct {
		Error *domain.Error `json:"error"`
	}{
		Error: err,
	})
}

func regexInt(idName string) string {
	return "/{" + idName + ":[0-9]+}"
}
