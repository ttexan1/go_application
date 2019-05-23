package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ttexan1/golang-simple/adapters/web"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
	"github.com/ttexan1/golang-simple/providers/sql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	envPath := os.Getenv("ENV_PATH")
	config, err := domain.NewConfig(envPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	s, err := sql.NewStorage(config.Driver, config.PostgresURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer s.Close()
	if config.IsDevelopment() {
		s.Migrate()
	}

	e := engine.NewEngine(s)

	log.Printf("Listening port %s ...\n", config.Port)
	http.ListenAndServe(config.Port, web.NewAdapter(e, config))
}
