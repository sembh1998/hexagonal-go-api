package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/courses"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/health"
	"github.com/sembh1998/hexagonal-go-api/kit/command"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	commandBus command.Bus
}

func New(host string, port int, commandBus command.Bus) Server {
	srv := Server{
		engine:     gin.New(),
		httpAddr:   fmt.Sprintf("%v:%v", host, port),
		commandBus: commandBus,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
