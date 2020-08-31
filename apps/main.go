package main

import (
	"ZebraX/apps/base/endpoint"
	"ZebraX/apps/config"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine = gin.New()
)

func init() {
	l, _ := os.Create("logger/logger__request.log")
	l.Chmod(644)

	s, _ := os.Create("logger/logger__request-recovery.log")
	s.Chmod(644)

	var (
		RECOVERY = gin.RecoveryWithWriter(io.MultiWriter(s, os.Stdout))
		LOGGER   = gin.LoggerWithWriter(io.MultiWriter(l, os.Stdout))
	)

	r.Use(RECOVERY, LOGGER)
}

func main() {
	log.Fatal(serverHandle().Run(config.APPLICATIONPORT))
}

func serverHandle() *gin.Engine {
	v1 := r.Group("/", endpoint.ValidationRequest)
	{
		v1.POST("/student", endpoint.InsertStudentEndpoint)
		v1.PUT("/student/:id", endpoint.PutStudentEndpoint)
		v1.GET("/student/:id", endpoint.GetStudentEndpoint)
		v1.DELETE("/student/:id", endpoint.DeleteStudentEndpoint)
	}

	return r
}
