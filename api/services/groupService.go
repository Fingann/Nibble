package services

import (
	"github.com/Fingann/notifyGame-api/database"
	"github.com/Fingann/notifyGame-api/models"
	"gorm.io/gorm"
)

type groupService struct {
	database *gorm.DB
}

func NewGroupService(database *gorm.DB) GroupService {
	return &groupService{
		database: database,
	}
}

func (gs *groupService) Database() *gorm.DB {
	return gs.database
}
