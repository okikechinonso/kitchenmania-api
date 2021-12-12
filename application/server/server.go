package server

import (
	"kitchenmaniaapi/application/handlers"
	"kitchenmaniaapi/infrastructure/persistence"
	"log"
	"net/http"
	"os"

	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB persistence.Database
	App handlers.App
}

func (s *Server) defineRoute(route *gin.Engine) {
	route.POST("/signup", s.App.SignUp())
}

func (s *Server) setupRoute() *gin.Engine {
	r := gin.New()

	s.defineRoute(r)
	return r
}

func (s *Server) Start() {
	r := s.setupRoute()
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable to start serve", err)
	}
}
