package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/amoako-franque/go-ecom-api/service/product"
	"github.com/amoako-franque/go-ecom-api/service/user"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Add routes for product
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	// Add routes for cart and orders
	// orderStore := order.NewStore(s.db)
	// cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	// cartHandler.RegisterRoutes(subrouter)

	// Serve static files
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Printf("Listening on port http://localhost%v", s.addr)

	return http.ListenAndServe(s.addr, router)
}
