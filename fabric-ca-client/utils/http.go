package utils

import "github.com/gin-gonic/gin"

func HttpSucess(status string, data interface{}) gin.H {
	return gin.H{
		"status": status,
		"data":   data,
	}
}

func HttpError(status string, meesage string) gin.H {
	return gin.H{
		"status": status,
		"error":  meesage,
	}
}
