// Package errors contains the error handler controller
package errors

import (
	"net/http"

	domainErrors "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/errors"
	"github.com/gin-gonic/gin"
)

type MessagesResponse struct {
	Message string `json:"message"`
}

func Handler(c *gin.Context) {
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domainErrors.AppError)
		if ok {
			resp := MessagesResponse{Message: err.Error()}
			switch err.Type {
			case domainErrors.NotFound:
				c.JSON(http.StatusNotFound, resp)
				return
			case domainErrors.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case domainErrors.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case domainErrors.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
				return
			case domainErrors.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
				return
			case domainErrors.RepositoryError:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			default:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			}
		}

		return
	}
}
