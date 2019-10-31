package net

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rapidvirt/host-agent/virt"
	log "github.com/sirupsen/logrus"
)

// Configuration to start the server
type Configuration struct {
	Address string
	Port    int
}

// ToString transforms the Configuration struct to a readable string
func (c *Configuration) ToString() string {
	return fmt.Sprintf("%s:%d", c.Address, c.Port)
}

// Server is the application main entry point controller
type Server struct {
	router   *mux.Router
	server   *http.Server
	virtConn *virt.Connection
}

// NewServer creates a new server struct with all default values
func NewServer() *Server {
	server := &Server{}
	server.router = mux.NewRouter()
	return server
}

// Initialize is the function to set up configuration into the server
func (s *Server) Initialize(conf *Configuration) {
	server := &http.Server{
		Addr:    conf.ToString(),
		Handler: s.router,
	}
	s.server = server
}

// Run this server
func (s *Server) Run() error {
	conn, err := virt.NewConnection()
	if err != nil {
		return err
	}
	s.virtConn = conn

	log.Infof("Server %v", s.server.Addr)
	return s.server.ListenAndServe()
}
