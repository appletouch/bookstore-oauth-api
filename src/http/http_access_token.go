package http

import (
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccesstokenHandlerInterface interface {
	GetById(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.ServiceInterface
}

func (handler *accessTokenHandler) GetById(ctx *gin.Context) {
	accesstokenId := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accesstokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, accessToken)

}

func NewAccessTokenHandler(service access_token.ServiceInterface) AccesstokenHandlerInterface {
	return &accessTokenHandler{
		service: service,
	}

}
