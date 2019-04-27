package web

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"github.com/ttexan1/golang-simple/domain"
)

type s struct{}

func (s *s) ReadDoc() string {
	f, err := domain.Assets.Open("swagger.yml")
	if err != nil {
		log.Print("Failed to open swagger documentation.")
		return ""
	}
	doc, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Failed to read swagger documentation.")
		return ""
	}
	return string(doc)
}

func init() {
	swag.Register(swag.Name, &s{})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger_index.html", http.StatusSeeOther)
}

func initDocs(m *mux.Router) {
	m.HandleFunc("/", redirect).Methods(http.MethodGet)
	m.PathPrefix("/").Handler(httpSwagger.WrapHandler)
}
