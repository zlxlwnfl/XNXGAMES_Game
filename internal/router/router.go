package router

import (
	"github.com/gin-gonic/gin"

	"XNXGAMES_Game/internal/handler"
)

func SetupRouter(gameHandler handler.GameHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1/games")
	{
		v1.POST("", gameHandler.CreateGame)
		v1.GET("/:id", gameHandler.GetGame)
		v1.GET("", gameHandler.GetAllGames)
		v1.PUT("/:id", gameHandler.UpdateGame)
		v1.DELETE("/:id", gameHandler.DeleteGame)
	}

	return router
}
