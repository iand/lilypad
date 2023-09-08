package solver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bacalhau-project/lilypad/pkg/server"
	"github.com/bacalhau-project/lilypad/pkg/system"
	"github.com/gorilla/mux"
)

type solverServer struct {
	options server.ServerOptions
}

func NewSolverServer(
	options server.ServerOptions,
	controller *SolverController,
) (*solverServer, error) {
	server := &solverServer{
		options: options,
	}
	return server, nil
}

func (solverServer *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.Use(server.CorsMiddleware)
	subrouter.Use(server.AuthMiddleware)

	subrouter.HandleFunc("/job_offers", server.Wrapper(solverServer.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", server.Wrapper(solverServer.addJobOffer)).Methods("POST")

	subrouter.HandleFunc("/resource_offers", server.Wrapper(solverServer.getResourceOffers)).Methods("GET")
	subrouter.HandleFunc("/resource_offers", server.Wrapper(solverServer.addResourceOffer)).Methods("POST")

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", solverServer.options.Host, solverServer.options.Port),
		WriteTimeout:      time.Minute * 15,
		ReadTimeout:       time.Minute * 15,
		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout:       time.Minute * 60,
		Handler:           router,
	}
	return srv.ListenAndServe()
}

func (solverServer *solverServer) getJobOffers(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}

func (solverServer *solverServer) getResourceOffers(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}

func (solverServer *solverServer) addJobOffer(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}

func (solverServer *solverServer) addResourceOffer(res http.ResponseWriter, req *http.Request) ([]string, error) {
	return []string{"job1", "job2"}, nil
}