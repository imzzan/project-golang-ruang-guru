package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseAPI interface {
	AddCourse(c *gin.Context)
}

type courseAPI struct {
	courseRepo repo.CourseRepository
}

func NewCourseAPI(courseRepo repo.CourseRepository) *courseAPI {
	return &courseAPI{courseRepo}
}

func (cr *courseAPI) AddCourse(c *gin.Context) {
	var course model.Course

	err := c.ShouldBindJSON(&course)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	err = cr.courseRepo.Store(&course)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "add course success",
	})
}
