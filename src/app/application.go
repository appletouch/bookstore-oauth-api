package app

import (
	"github.com/appletouch/bookstore-oauth-api/src/http"
	"github.com/appletouch/bookstore-oauth-api/src/repository/db"
	"github.com/appletouch/bookstore-oauth-api/src/repository/rest"
	"github.com/appletouch/bookstore-oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	ginEngine = gin.Default()
)

func StartApplication() {

	atHandler := http.NewAccessTokenHandler(services.NewService(rest.NewRespository(), db.NewRepository()))

	ginEngine.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	ginEngine.POST("/oauth/access_token", atHandler.Create)

	ginEngine.Run(":3001")

}
