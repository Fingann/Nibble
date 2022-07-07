package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name    string
	OwnerID uint
	Owner   User     `gorm:"foreignkey:OwnerID;referances:ID;constraint:fk_groups_owner"`
	Games   []Game   `gorm:"many2many:group_games"`
	Members []Member `gorm:"many2many:group_members"`
}

type Member struct {
	gorm.Model
	GroupId uint
	UserId  uint
}

type Game struct {
	gorm.Model
	Name string
	Logo []byte
}

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
	Groups       []Group `gorm:"many2many:group_members"`
}
