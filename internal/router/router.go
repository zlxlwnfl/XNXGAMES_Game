package router

import (
	"github.com/gin-gonic/gin"

	"XNXGAMES_Game/internal/handler"
)

func SetupGameRouter(gh handler.GameHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/games", gh.CreateGame)
		v1.GET("/games/:id", gh.GetGame)
		v1.GET("/games", gh.GetAllGames)
		v1.PUT("/games/:id", gh.UpdateGame)
		v1.DELETE("/games/:id", gh.DeleteGame)
	}

	return router
}
