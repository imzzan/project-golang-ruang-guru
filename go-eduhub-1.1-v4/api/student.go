package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	GetStudents(c *gin.Context)
	GetStudentByID(c *gin.Context)
}

type studentAPI struct {
	studentRepo repo.StudentRepository
}

func NewStudentAPI(studentRepo repo.StudentRepository) *studentAPI {
	return &studentAPI{studentRepo}
}

func (s *studentAPI) AddStudent(c *gin.Context) {
	var students model.Student

	err := c.ShouldBindJSON(&students)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	err = s.studentRepo.Store(&students)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "add student success",
	})
}

func (s *studentAPI) GetStudents(c *gin.Context) {
	students, err := s.studentRepo.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (s *studentAPI) GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	students, err := s.studentRepo.FetchAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	for _, value := range students {
		if value.ID == id {
			c.JSON(http.StatusOK, value)
			return
		}
	}

	c.AbortWithStatusJSON(404, model.ErrorResponse{
		Error: "Data Not Found",
	})
}
