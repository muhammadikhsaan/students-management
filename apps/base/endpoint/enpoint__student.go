package endpoint

import (
	"ZebraX/apps/base/model"
	"ZebraX/apps/base/repository"
	"ZebraX/apps/config"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var student repository.StudentsInterface

func init() {
	student = repository.NewStudentRepository(context.Background(), config.Getdatabase())
}

//ValidationRequest used to validation reguest header
func ValidationRequest(c *gin.Context) {
	var h model.HeaderModel

	c.Header("Accept", "application/json")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")

	if err := c.ShouldBindHeader(&h); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

//InsertStudentEndpoint to handle request to add students
func InsertStudentEndpoint(c *gin.Context) {
	var s model.StudentModel

	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := student.InsertStudent(&s)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{})
		return
	}
}
