package creating

import (
	"context"

	mooc "github.com/sembh1998/hexagonal-go-api/internal"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

func (s CourseService) CreateCourse(ctx context.Context, ID, Name, Duration string) error {
	course, err := mooc.NewCourse(ID, Name, Duration)
	if err != nil {
		return err
	}
	return s.courseRepository.Save(ctx, course)
}
