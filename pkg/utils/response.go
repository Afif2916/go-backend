package utils 

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}


func Success(c *gin.Context, code int, message string, data interface{}) {
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(code, response)
}

func Error(c *gin.Context, code int, message string) {
	response := Response{
		Status: "error",
		Error:  message,
	}
	c.JSON(code, response)
}