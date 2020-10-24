package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/emadghaffari/grpc_oauth_client/grpc_client/controllers"
	"github.com/gin-gonic/gin"
)

var(
	// AccessTokenMiddleware handler
	AccessTokenMiddleware AccessTokenMiddlewareInterface = &accessTokenMiddleware{}
)

// AccessTokenMiddlewareInterface interface
type AccessTokenMiddlewareInterface interface{
	CheckMiddleware(c *gin.Context) 
}

type accessTokenMiddleware struct{}

func (ac *accessTokenMiddleware) CheckMiddleware(c *gin.Context) {
	token := strings.TrimSpace(c.Request.Header.Get("api_token"))
	if token == ""{
		Middleware.RespondWithErrorJSON(c,http.StatusBadRequest,"the access token is null.")
		return
	}
	result, err := controllers.ClientAccessToken.Get(token)
	if err !=nil{
		Middleware.RespondWithErrorJSON(c,err.Status(),err.Message())
		return
	}
	c.Request.Header.Add("client_id",strconv.Itoa(int(result.ClientId)))
	c.Request.Header.Add("user_id",strconv.Itoa(int(result.UserId)))
	c.Next()
}