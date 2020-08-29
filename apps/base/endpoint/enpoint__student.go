package endpoint

import (
	"ZebraX/apps/base/model"
	"ZebraX/apps/base/repository"
	"ZebraX/apps/config"
	"ZebraX/apps/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	student      repository.StudentsInterface
	authorizaton = services.NewLoginService()
)

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

	if !authorizaton.HandleAuthenticator(h.Authorization) {
		c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
		return
	}
}

//InsertStudentEndpoint to handle request POST to add students
func InsertStudentEndpoint(c *gin.Context) {
	var request model.StudentModel

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := student.InsertStudent(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Insert Student"})
}

//PutStudentEndpoint to handle request PUT to change/update students
func PutStudentEndpoint(c *gin.Context) {
	var request model.StudentModel

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request.ID = id
	_, err = student.UpdateStudent(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Update Student"})
}
