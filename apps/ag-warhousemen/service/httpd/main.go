package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	Setup()

	engine := gin.Default()
	hDependencies := NewHandlerDependency()
	InitRoutes(engine, hDependencies)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
