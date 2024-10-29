package server

import (
	"fmt"
	"net/http"
	"github.com/your_project/utils"
)

// Server is a struct that encapsulates server functionalities.
type Server struct {
	host   string
	port   int
	logger *utils.Logger
}

// NewServer initializes and returns a new Server instance with default settings.
func NewServer(host string, port int, logger *utils.Logger) *Server {
	return &Server{
		host:   host,
		port:   port,
		logger: logger,
	}
}

// StartServer starts the HTTP server and listens for incoming requests.
func (s *Server) StartServer() error {
	address := fmt.Sprintf("%s:%d", s.host, s.port)
	s.logger.LogInfo(fmt.Sprintf("Starting server at %s", address))
	http.HandleFunc("/", s.HandleRequests)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		s.logger.LogError(fmt.Sprintf("Failed to start server: %v", err))
		return err
	}
	return nil
}

// HandleRequests handles incoming HTTP requests.
func (s *Server) HandleRequests(w http.ResponseWriter, r *http.Request) {
	s.logger.LogInfo("Received request")
	// Here you can add logic to handle different routes and methods
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request handled successfully"))
}

// DispatchTask dispatches a task based on the taskID.
func (s *Server) DispatchTask(taskID string) error {
	s.logger.LogInfo(fmt.Sprintf("Dispatching task with ID: %s", taskID))
	// Implement task dispatching logic here
	// For now, we just log the task dispatch
	return nil
}
