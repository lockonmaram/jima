package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	StatusMessageSuccess        = "success"
	StatusMessageError          = "error"
	StatusMessageInvalidRequest = "invalid request"

	HeaderAuthorization = "Authorization"

	ContextUserAuth = "userAuth"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
	Error   string `json:"error,omitempty"`
}

func HandleResponse(c *gin.Context, resp Response) {
	responseMessage := StatusMessageSuccess

	if resp.Status == 400 {
		responseMessage = StatusMessageInvalidRequest
	} else if resp.Status > 400 {
		responseMessage = StatusMessageError
	}

	if resp.Message != "" {
		responseMessage = resp.Message
	}

	response := Response{
		Status:  resp.Status,
		Message: responseMessage,
		Data:    resp.Data,
		Meta:    resp.Meta,
		Error:   resp.Error,
	}

	c.JSON(resp.Status, response)
}

func HandleRequest(c *gin.Context, request any) (err error) {
	if err := c.BindJSON(request); err != nil {
		HandleResponse(c, Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return err
	}

	if err := ValidateStruct(request); err != nil {
		HandleResponse(c, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   ErrInvalidRequest.Error(),
		})
		return err
	}

	return nil
}

func GetUserAuthClaims(c *gin.Context) *Claims {
	userAuth, exists := c.MustGet(ContextUserAuth).(*Claims)
	if !exists {
		return nil
	}

	return userAuth
}
