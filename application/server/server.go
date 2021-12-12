package server

import (
	"kitchenmaniaapi/application/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


type Server struct{
	App handlers.App
}

func(s *Server) defineRoute(route *gin.Engine){
	route.POST("/signup", s.App.SignUp())
}


func (s *Server) setupRoute() *gin.Engine{
	r := gin.New()
	s.defineRoute(r)
	return r
}

func (s *Server) Start(){
	r := s.setupRoute()
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr: port,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil{
		log.Fatal("unable to start serve")
	}
}