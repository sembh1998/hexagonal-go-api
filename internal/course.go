package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

var ErrInvalidCourseID = errors.New("invalid Course ID")

type CourseID struct {
	value string
}

func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}
	return CourseID{
		value: v.String(),
	}, nil
}

var ErrInvalidCourseName = errors.New("invalid Course Name")

type CourseName struct {
	value string
}

func NewCourseName(value string) (CourseName, error) {
	if len(value) == 0 {
		return CourseName{}, fmt.Errorf("%w: %s", ErrInvalidCourseName, value)
	}
	return CourseName{
		value: value,
	}, nil
}

var ErrInvalidCourseDuration = errors.New("invalid Course Duration")

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if len(value) == 0 {
		return CourseDuration{}, fmt.Errorf("%w: %s", ErrInvalidCourseDuration, value)
	}
	return CourseDuration{
		value: value,
	}, nil
}

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

func (c Course) ID() string {
	return c.id.value
}

func (c Course) Name() string {
	return c.name.value
}

func (c Course) Duration() string {
	return c.duration.value
}
