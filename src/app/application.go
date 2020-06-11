package app

import (
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/appletouch/bookstore-oauth-api/src/http"
	"github.com/appletouch/bookstore-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	ginEngine = gin.Default()
)

func StartApplication() {

	accesstokenHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	ginEngine.GET("/oauth/access_token/:access_token_id", accesstokenHandler.GetById)
	ginEngine.POST("/oauth/access_token", accesstokenHandler.Create)

	ginEngine.Run(":3001")

}
