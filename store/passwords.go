package store

import (
	"time"

	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	ID                int    `json:"password_id" gorm:"primaryKey"`
	Name              string `gorm:"uniqueIndex"`
	Url               string `gorm:"uniqueIndex"`
	Username          string
	EncryptedPassword []byte `gorm:"type:blob"`
	IV                []byte `gorm:"type:blob"`
	Notes             string
	LastUsed          time.Time
	User              string
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
	var names []string
	if err := db.DB.Model(&Password{}).Select("name").Scan(&names).Error; err != nil {
		return nil, err
	}
	return names, nil
}

func (db *Database) DeletePassword(password *Password) error {
	return db.DB.Delete(password).Error
}

func (db *Database) UpdatePassword(password *Password) error {
	return db.DB.Save(password).Error
}
