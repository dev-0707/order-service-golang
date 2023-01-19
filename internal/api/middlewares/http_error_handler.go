package middlewares

import (
	"fmt"
	"net/http"
	"order-service/internal/api/problems"
	"order-service/internal/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func HandleValidationError(c *gin.Context, err error) {
	// var logMessage, statusCode, message, errorType = buildError(c, err)
	var problem, logMessage = buildError(c, err)
	log.Error().Err(err).Msg(logMessage)
	c.JSON(http.StatusInternalServerError, problem)
	return
}

func buildError(c *gin.Context, err error) (*problems.Problem, string) {
	var logMessage string
	var statusCode int
	var message string
	var errorType string
	if ok := err.(*service.MyError); ok != nil {
		logMessage = fmt.Sprintf("Validation error '%d'", c.Request.URL.Path)
		errorType = "Validation error"
		statusCode = http.StatusBadRequest
		message = "Error validating request"
	} else {
		logMessage = fmt.Sprintf("Interal server error '%d'", c.Request.URL.Path)
		errorType = "Internal server error"
		statusCode = http.StatusInternalServerError
		message = "An interal server error has occurred"
	}
	problem := problems.NewProblem(statusCode, c.Request.URL.Path, errorType, message)
	return &problem, logMessage
}

// NoMethodHandler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Metodo no permitido"})
	}
}

// NoRouteHandler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	}
}
