package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mooc "github.com/sembh1998/hexagonal-go-api/internal"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/courses"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	courseRepository mooc.CourseRepository
}

func New(host string, port int, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%v:%v", host, port),
		courseRepository: courseRepository,
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
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
