package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"XNXGAMES_Game/internal/dto"
	"XNXGAMES_Game/internal/entity"
	"XNXGAMES_Game/internal/service"
)

type GameHandler struct {
	gameService service.GameService
}

func NewGameHandler(gs service.GameService) GameHandler {
	return GameHandler{gameService: gs}
}

func (gh GameHandler) CreateGame(ctx *gin.Context) {
	var game entity.Game

	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := gh.gameService.CreateGame(ctx, game)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, game)
}

func (gh GameHandler) GetGame(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효한 id가 아닙니다."})
		return
	}

	game, err := gh.gameService.GetGame(ctx, id)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (gh GameHandler) GetAllGames(ctx *gin.Context) {
	queryOffset := ctx.DefaultQuery("offset", "0")
	queryLimit := ctx.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(queryOffset)
	if err != nil {
		return
	}

	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		return
	}

	req := dto.GetAllGamesRequest{
		Offset: offset,
		Limit:  limit,
	}

	game, err := gh.gameService.GetAllGames(ctx, req)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (gh GameHandler) UpdateGame(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효한 id가 아닙니다."})
		return
	}

	var req dto.UpdateGameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, err := gh.gameService.UpdateGame(ctx, id, req)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, game)
}

func (gh GameHandler) DeleteGame(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "유효한 id가 아닙니다."})
		return
	}

	err = gh.gameService.DeleteGame(ctx, id)
	if err != nil {
		return
	}

	ctx.Status(http.StatusOK)
}
