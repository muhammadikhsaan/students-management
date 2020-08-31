package main

import (
	"ZebraX/base/endpoint"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidationRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, err := http.NewRequest("GET", "http://localhost:8080/student/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")
	c.Request = request
	endpoint.ValidationRequest(c)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testArray := []struct {
		Name string
		URL  string
		body []byte
	}{
		{Name: "data available", URL: "http://localhost:8080/student/1"},
		{Name: "data not available", URL: "http://localhost:8080/student/10"},
	}

	for _, testValue := range testArray {
		t.Run(testValue.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			_, r := gin.CreateTestContext(response)

			request, err := http.NewRequest(http.MethodGet, testValue.URL, nil)
			if err != nil {
				t.Fatalf("could not get request: %v", err)
			}

			request.Header.Set("Authorization", "Bearer $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")

			v1 := r.Group("/", endpoint.ValidationRequest)
			{
				v1.GET("/student/:id", endpoint.GetStudentEndpoint)
			}

			r.ServeHTTP(response, request)
			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}

func TestPostHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testArray := []struct {
		Name string
		URL  string
		body []byte
	}{
		{Name: "data not ready", URL: "http://localhost:8080/student", body: []byte(`{"id":2, "name":"Muhammad Ikhsan", "age":1}`)},
		{Name: "data ready", URL: "http://localhost:8080/student", body: []byte(`{"id":2, "name":"Muhammad Ikhsan", "age":1}`)},
	}

	for _, testValue := range testArray {
		t.Run(testValue.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			_, r := gin.CreateTestContext(response)

			request, err := http.NewRequest(http.MethodPost, testValue.URL, bytes.NewBuffer(testValue.body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			request.Header.Set("Authorization", "Bearer $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")

			v1 := r.Group("/", endpoint.ValidationRequest)
			{
				v1.POST("/student", endpoint.InsertStudentEndpoint)
			}

			r.ServeHTTP(response, request)
			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}

func TestPutHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testArray := []struct {
		Name string
		URL  string
		body []byte
	}{
		{Name: "data not ready", URL: "http://localhost:8080/student/1", body: []byte(`{"name":"Muhammad Ikhsan", "age":5}`)},
		{Name: "data readt", URL: "http://localhost:8080/student/4", body: []byte(`{"name":"Muhammad Ikhsan", "age":20}`)},
	}

	for _, testValue := range testArray {
		t.Run(testValue.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			_, r := gin.CreateTestContext(response)

			request, err := http.NewRequest(http.MethodPut, testValue.URL, bytes.NewBuffer(testValue.body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			request.Header.Set("Authorization", "Bearer $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")

			v1 := r.Group("/", endpoint.ValidationRequest)
			{
				v1.PUT("/student/:id", endpoint.PutStudentEndpoint)
			}

			r.ServeHTTP(response, request)
			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}

func TestDeleteHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testArray := []struct {
		Name string
		URL  string
		body []byte
	}{
		{Name: "data not ready", URL: "http://localhost:8080/student/1"},
		{Name: "data ready", URL: "http://localhost:8080/student/3"},
	}

	for _, testValue := range testArray {
		t.Run(testValue.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			_, r := gin.CreateTestContext(response)

			request, err := http.NewRequest(http.MethodDelete, testValue.URL, nil)
			if err != nil {
				assert.Error(t, err)
			}

			request.Header.Set("Authorization", "Bearer $2a$08$1rzzWv6InEiwn8nDAnK0iuOdl6v0ySdYr8jir5mWYzglDmkXypr9W")

			v1 := r.Group("/", endpoint.ValidationRequest)
			{
				v1.DELETE("/student/:id", endpoint.DeleteStudentEndpoint)
			}

			r.ServeHTTP(response, request)
			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}
