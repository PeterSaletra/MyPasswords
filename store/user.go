package store

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `json:"user_id" gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex"`
	Password  []byte    `gorm:"type:blob"`
	LastLogin time.Time `json:"last_login" gorm:"autoCreateTime"`

	Passwords []Password `gorm:"foreignKey:UserID;references:ID"`
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

func (db *Database) UpdateLastLogin(user *User) error {
	return db.DB.Model(user).Update("last_login", time.Now()).Error
}
