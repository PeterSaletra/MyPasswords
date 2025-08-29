package store

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"uniqueIndex"`
	LastLogin time.Time `json:"last_login" gorm:"autoCreateTime"`

	Passwords []Password `gorm:"foreignKey:UserID;references:Username"`
}

func (db *Database) CreateUser(user *User) error {
	return db.DB.Create(user).Error
}

func (db *Database) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *Database) UpdateLastLogin(username string) error {
	return db.DB.Model(&User{}).Where("username = ?", username).Update("last_login", time.Now()).Error
}
