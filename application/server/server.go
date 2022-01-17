package server

import (
	"kitchenmaniaapi/application/handlers"
	"kitchenmaniaapi/application/middleware"
	"kitchenmaniaapi/infrastructure/persistence"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	DB persistence.Database
	App handlers.App
}

func (s *Server) defineRoute(router *gin.Engine) {
	apirouter := router.Group("/api/v1")
	apirouter.POST("/signup", s.App.SignUp())
	apirouter.POST("/login",s.App.Login())

	authorized := apirouter.Group("/")
	authorized.Use(middleware.Authorize(s.DB.FindUserByEmail,s.DB.TokenInBlacklist))
	authorized.POST("/create",s.App.CreateBlog())
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
