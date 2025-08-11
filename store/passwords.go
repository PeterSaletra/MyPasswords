package store

import (
	"time"

	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	ID                int    `json:"password_id" gorm:"primaryKey"`
	Name              string `gorm:"uniqueIndex"`
	Website           string `gorm:"uniqueIndex"`
	Username          string
	EncryptedPassword []byte `gorm:"type:blob"`
	VI                []byte `gorm:"type:blob"`
	Notes             string
	LastUsed          time.Time
}

func (db *Database) CreatePassword(password *Password) error {
	return db.DB.Create(password).Error
}

func (db *Database) GetPasswordByName(name string) (*Password, error) {
	var password Password
	if err := db.DB.Where("name = ?", name).First(&password).Error; err != nil {
		return nil, err
	}
	return &password, nil
}

func (db *Database) GetAllPasswordsNames() ([]string, error) {
	var passwords []Password
	if err := db.DB.Find(&passwords).Error; err != nil {
		return nil, err
	}
	var names []string
	for _, p := range passwords {
		names = append(names, p.Name)
	}
	return names, nil
}

func (db *Database) DeletePassword(password *Password) error {
	return db.DB.Delete(password).Error
}

func (db *Database) UpdatePassword(password *Password) error {
	return db.DB.Save(password).Error
}
