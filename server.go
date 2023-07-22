package sharedtimer

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type ServerConfig struct {
	PORT string
}

type server struct {
	cnf    ServerConfig
	router chi.Router
	timers map[string]*Timer
}

func NewServer(cnf ServerConfig) server {
	s := server{
		cnf:    cnf,
		router: chi.NewRouter(),
		timers: map[string]*Timer{},
	}

	s.router.Mount("/", s.Routes())

	// TODO: load all timers from datastore into memory
	// TODO: run all timers asynchronously if they weren't paused

	return s
}

func (svr *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svr.router.ServeHTTP(w, r)
}

func (svr *server) Run() error {
	fmt.Printf("running server on port %s\n", svr.cnf.PORT)
	return http.ListenAndServe(":"+svr.cnf.PORT, svr)
}

func (svr *server) OnShutdown() {
	// TODO: store all timers in datastore
}
