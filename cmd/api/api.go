package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restapi-go/services/users"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
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
	userStore := users.NewUserStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)
	println("Starting API server on", s.addr)

	// return http.ListenAndServe(s.Addr, corsHandler)
	return http.ListenAndServe(s.addr, router)
}
