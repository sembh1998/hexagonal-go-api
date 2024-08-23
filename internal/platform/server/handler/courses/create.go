package courses

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/sembh1998/hexagonal-go-api/internal"
	"github.com/sembh1998/hexagonal-go-api/internal/creating"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(creatingCourseService creating.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := creatingCourseService.CreateCourse(ctx, req.ID, req.Name, req.Duration); err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID), errors.Is(err, mooc.ErrInvalidCourseName), errors.Is(err, mooc.ErrInvalidCourseDuration):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
