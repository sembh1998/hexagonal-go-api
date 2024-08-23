package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sembh1998/hexagonal-go-api/internal/creating"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/server"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "train"
	dbPass = "train"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "train"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	courseService := creating.NewCourseService(courseRepository)

	srv := server.New(host, port, courseService)
	return srv.Run()
}
