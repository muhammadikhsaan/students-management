package main

import (
	"ZebraX/apps/base/endpoint"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func PostRequestTest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, err := http.NewRequest("GET", "http://localhost:8080/student/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	request.Header.Set("Authorization", "$2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")
	c.Request = request
	endpoint.GetStudentEndpoint(c)

	assert.Equal(t, http.StatusOK, response.Code)
}
