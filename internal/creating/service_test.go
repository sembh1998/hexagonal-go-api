package creating

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/sembh1998/hexagonal-go-api/internal"
	"github.com/sembh1998/hexagonal-go-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_CreateCourse_RespositoryError(t *testing.T) {
	courseID := "01917342-4632-7bdb-9f64-2d700bd38813"
	courseName := "Test del Curso"
	courseDuration := "1 hora"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)
	require.NotEmpty(t, course)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeded(t *testing.T) {
	courseID := "01917342-4632-7bdb-9f64-2d700bd38813"
	courseName := "Test del Curso"
	courseDuration := "1 hora"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)
	require.NotEmpty(t, course)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
