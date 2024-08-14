package mysql

import (
	"database/sql"

	mooc "github.com/sembh1998/hexagonal-go-api/internal"
)

func NewCourseRepository(db *sql.DB) mooc.CourseRepository {
	return nil
}
