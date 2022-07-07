package services

import (
	"github.com/Fingann/notifyGame-api/database"
	"github.com/Fingann/notifyGame-api/models"
	"gorm.io/gorm"
)

type gameService struct {
	database *gorm.DB
}

func NewGameService(database *gorm.DB) GameService {
	return &gameService{
		database: database,
	}
}

func (gs *gameService) Database() *gorm.DB {
	return gs.database
}
