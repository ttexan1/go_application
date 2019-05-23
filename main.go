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
	if config.IsDevelopment() || config.IsSandbox() {
		s.Migrate()
	}

	e := engine.NewEngine(s)

	log.Printf("Listening port %s ...\n", ":9000")
	http.ListenAndServe("127.0.0.1:9000", web.NewAdapter(e, config))
}
