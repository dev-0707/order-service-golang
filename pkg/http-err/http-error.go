package http_err

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

type HttpResponse struct {
	Type     string
	Title    string
	Status   int
	Detail   string
	Instance string
}

func InternalServerErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	// httpResponse := HttpResponse{Message: "Si è verificato un errore", Status: 500}
	httpResponse := HttpResponse{Detail: "Si è verificato un errore", Status: 500, Type: goErr.Error()}
	c.AbortWithStatusJSON(500, httpResponse)
}
