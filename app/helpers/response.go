package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func ResponseJSON(c *gin.Context, code int, message string, data interface{}) {
	status := code >= 200 && code < 300
	c.JSON(code, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func ValidationErrors(err error) interface{} {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = fieldError.Tag()
		}
		return errors
	}
	return nil
}