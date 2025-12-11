package helper

import (
	"errors"
	"io"
	"jima/entity/model"
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

	if resp.Status == 0 {
		resp.Status = http.StatusOK
	} else if resp.Status == 400 {
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
	if err := c.BindUri(request); err != nil {
		HandleResponse(c, Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return err
	}

	if err := c.BindQuery(request); err != nil {
		HandleResponse(c, Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return err
	}

	if err := c.ShouldBindJSON(request); err != nil && !errors.Is(err, io.EOF) {
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

func IsUserAdminOrSelf(c *gin.Context, serial string) bool {
	userAuth := GetUserAuthClaims(c)
	if userAuth.Role != string(model.UserRoleAdmin) && userAuth.Serial != serial {
		return false
	}
	return true
}

func IsUserGroupManagerOrSelf(userGroup *model.UserGroup, userSerial string) bool {
	return userGroup.Role == model.UserGroupRoleManager || userGroup.UserSerial == userSerial
}
