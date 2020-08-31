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

//InsertStudentEndpoint to handle request POST to add student
func InsertStudentEndpoint(c *gin.Context) {
	var request model.StudentModel

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := student.InsertStudent(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	_, errs := result.RowsAffected()

	if errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errs.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Insert Student"})
}

//PutStudentEndpoint to handle request PUT to change/update student
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
	result, errs := student.UpdateStudent(&request)

	if errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errs.Error()})
		return
	}

	row, errss := result.RowsAffected()

	if errss != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errss.Error()})
		return
	}

	if row == 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Failed to Update Data"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Update Student with ID " + c.Param("id")})
}

//GetStudentEndpoint to handle request GET to select student
func GetStudentEndpoint(c *gin.Context) {
	var request model.StudentModel

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request.ID = id
	result, errs := student.SelectStudent(&request)

	if errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errs.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, result)
}

//DeleteStudentEndpoint to handle request delete to delete student
func DeleteStudentEndpoint(c *gin.Context) {
	var request model.StudentModel

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	request.ID = id
	result, errs := student.DeleteStudent(&request)

	if errs != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errs.Error()})
		return
	}

	row, errss := result.RowsAffected()

	if errss != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": errss.Error()})
		return
	}

	if row == 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Failed to Delete Data"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Delete Student with ID " + c.Param("id")})
}
