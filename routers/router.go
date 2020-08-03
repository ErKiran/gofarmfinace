package routers

import (
	v1 "farmfinance/routers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.RedirectFixedPath = true
	r.RedirectTrailingSlash = true
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	{
		api.GET("health", func(c *gin.Context) { c.String(http.StatusOK, "Ok") })

		api.POST("farm",v1.CreateFarm)
	}
	return r
}
