package bootstrap

import "github.com/sembh1998/hexagonal-go-api/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
