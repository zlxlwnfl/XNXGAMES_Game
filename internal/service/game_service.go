package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"XNXGAMES_Game/internal/dto"
	"XNXGAMES_Game/internal/entity"
	"XNXGAMES_Game/internal/repository"
)

type GameService struct {
	gameRepo repository.GameRepository
}

func NewGameService(gr repository.GameRepository) GameService {
	return GameService{gameRepo: gr}
}

func (gc GameService) CreateGame(ctx *gin.Context, game entity.Game) (entity.Game, error) {
	if err := gc.gameRepo.Create(game); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return game, err
	}

	return game, nil
}

func (gc GameService) GetGame(ctx *gin.Context, id int) (entity.Game, error) {
	game, err := gc.gameRepo.Get(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return game, err
	}

	return game, nil
}

func (gc GameService) GetAllGames(c *gin.Context, req dto.GetAllGamesRequest) ([]entity.Game, error) {
	games, err := gc.gameRepo.GetAll(req.Offset, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return games, err
	}

	return games, nil
}

func (gc GameService) UpdateGame(ctx *gin.Context, id int, req dto.UpdateGameRequest) (entity.Game, error) {
	game, err := gc.gameRepo.Get(uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return game, err
	}

	game.Title = req.Title

	if err := gc.gameRepo.Update(game); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return game, err
	}

	return game, nil
}

func (gc GameService) DeleteGame(ctx *gin.Context, id int) error {
	if err := gc.gameRepo.Delete(uint64(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	return nil
}
