package http

// Handler dus not contain any business logic, but focuses on request processing and response creation.

import (
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/appletouch/bookstore-oauth-api/src/services"
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccesstokenHandlerInterface interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
type accessTokenHandler struct {
	service services.ServiceInterface
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context) {
	accesstokenId := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accesstokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, accessToken)

}

func (handler *accessTokenHandler) Create(ctx *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		restErr := errors.RestErr{400, "Bad Request", "Invalid Json body"}
		ctx.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := handler.service.Create(request)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)

}

//func (handler *accessTokenHandler) Create(ctx *gin.Context) {
//	var at access_token.AccessToken
//	if err := ctx.ShouldBindJSON(&at); err != nil {
//		restErr := errors.New(400, err.Error())
//		ctx.JSON(restErr.Status, restErr)
//		return
//	}
//	if err := handler.service.Create(at); err != nil {
//		ctx.JSON(err.Status, err)
//	}
//
//	ctx.JSON(http.StatusCreated, at)
//
//}

func NewAccessTokenHandler(service services.ServiceInterface) AccesstokenHandlerInterface {
	return &accessTokenHandler{
		service: service,
	}

}
