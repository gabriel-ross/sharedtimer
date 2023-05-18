package stgo

import (
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	PORT string
}

type server struct {
	cnf    ServerConfig
	router chi.Router
	timers map[uuid.UUID]*countdownTimer
}

func NewServer(cnf ServerConfig) server {
	return server{
		cnf:    cnf,
		router: chi.NewRouter(),
		timers: map[uuid.UUID]*countdownTimer{},
	}
}

func (svr *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svr.router.ServeHTTP(w, r)
}

func (svr *server) Run() error {
	return http.ListenAndServe(":"+svr.cnf.PORT, svr)
}

func foo() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	svr := grpc.NewServer()
	if err := svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
