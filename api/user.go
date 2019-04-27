package api

import (
	"fmt"
	"net/http"
)

// UserHandler is the handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is user api")
}

// ProductHandler is the handler
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is product api")
}

// TeacherHandler is the handler
func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is teacher api")
}
