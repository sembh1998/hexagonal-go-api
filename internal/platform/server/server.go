package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/hexagonal-go-api/internal/creating"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/courses"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	creatingCourseService creating.CourseService
}

func New(host string, port int, creatingCourseService creating.CourseService) Server {
	srv := Server{
		engine:                gin.New(),
		httpAddr:              fmt.Sprintf("%v:%v", host, port),
		creatingCourseService: creatingCourseService,
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
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
}
