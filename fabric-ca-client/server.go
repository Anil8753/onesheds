package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/anil8753/fabric-ca-client/handlers"
)

func StartServer() {

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := Setup()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/enrolladmin", handlers.EnrollAdminHandler())
		v1.POST("/registeruser", handlers.RegisterUserHandler())
		v1.POST("/revoke", handlers.RevokeUserHandler())

		v1.GET("/users/:id", handlers.UserHandler())
		v1.GET("/users", handlers.AllUsersHandler())
	}

	r.Run()
}

func Setup() *gin.Engine {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	return r
}
