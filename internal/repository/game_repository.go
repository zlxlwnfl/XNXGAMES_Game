package repository

import (
	"gorm.io/gorm"

	"XNXGAMES_Game/internal/entity"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return GameRepository{db: db}
}

func (r GameRepository) Create(game entity.Game) error {
	return r.db.Create(&game).Error
}

func (r GameRepository) Get(id uint64) (entity.Game, error) {
	var game entity.Game
	err := r.db.First(&game, id).Error
	return game, err
}

func (r GameRepository) GetAll(offset int, limit int) ([]entity.Game, error) {
	var games []entity.Game
	err := r.db.Offset(offset).Limit(limit).Find(&games).Error
	return games, err
}

func (r GameRepository) Update(game entity.Game) error {
	return r.db.Save(&game).Error
}

func (r GameRepository) Delete(id uint64) error {
	return r.db.Delete(&entity.Game{}, id).Error
}
