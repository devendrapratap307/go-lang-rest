package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		Addr: addr,
	}
}
func (s *APIServer) Run() error {
	// Start the API server
	// This is a placeholder implementation

	// create router
	router := mux.NewRouter()

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:  []string{"Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// })
	// corsHandler := c.Handler(router)

	// register services
	println("Starting API server on", s.Addr)

	// return http.ListenAndServe(s.Addr, corsHandler)
	return http.ListenAndServe(s.Addr, router)
}
