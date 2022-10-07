package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/vmw-pso/logger-service/data"
	"github.com/vmw-pso/toolkit"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 80, "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf(":%d", *port)

	srv := newServer()
	fmt.Printf("Starting authentication-service, listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	mux    *chi.Mux
	models *data.Models
	tools  toolkit.Tools
}

func newServer() *server {
	mux := chi.NewMux()
	models := data.New()
	tools := toolkit.Tools{}

	srv := &server{
		mux:    mux,
		models: models,
		tools:  tools,
	}
	srv.routes()

	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
