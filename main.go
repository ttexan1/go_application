package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"

	"github.com/ttexan1/golang-simple/adapters/web"
	"github.com/ttexan1/golang-simple/domain"
	"github.com/ttexan1/golang-simple/engine"
	"github.com/ttexan1/golang-simple/providers/psql"
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
	s, err := psql.NewStorage(config.Driver, config.PostgresURL)
	if err != nil {
		fmt.Println("aaa")
		log.Fatal(err.Error())
	}
	defer s.Close()
	if config.IsDevelopment() || config.IsSandbox() {
		s.Migrate()
	}

	e := engine.NewEngine(s)
	_ = e

	// env := "devlopment"
	env := "production"
	log.Printf("Listening port %s ...\n", ":9000")
	if env == "devlopment" {
		http.ListenAndServe("127.0.0.1:9000", web.NewAdapter(e, config))
	} else {
		l, err := net.Listen("tcp", "127.0.0.1:9000")
		if err != nil {
			return
		}
		fcgi.Serve(l, web.NewAdapter(e, config))
	}
}
