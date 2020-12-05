package service

import (
    "github.com/gorilla/mux"
    "github.com/rosish/inventory/InventoryAccess/internal/urls"
    "log"
    "net/http"
)

type Server struct {
    Router            *mux.Router
}


func NewServer() *Server {
    s := &Server{}
    return s
}

func (s *Server) StartServer() {
    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", s.Router))
}

func (s *Server) SetupRoutes() {
    s.Router = mux.NewRouter()
    urlPatterns := urls.ReturnURLS()
    s.Router.HandleFunc(urlPatterns.FETCH_INV_PATH, InventoryAccessHandler)
}