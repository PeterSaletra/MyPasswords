package store

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect(master_key string) error {
	homedDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	dbDir := filepath.Join(homedDir, ".local", "mypasswords", "db")
	dbFile := filepath.Join(dbDir, "mypasswords.db")

	if err := os.MkdirAll(dbDir, 0700); err != nil {
		return fmt.Errorf("failed to create database directory: %w", err)
	}

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.OpenFile(dbFile, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return fmt.Errorf("failed to create database file: %w", err)
		}
		file.Close()
	}

	dsn := fmt.Sprintf("file:%s?_pragma_key=%s&_pragma_cipher_page_size=%d", dbFile, master_key, 4096)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now()
		},
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	d.DB = db
	return nil
}

func (d *Database) Migrate() error {
	err := d.DB.AutoMigrate(&User{}, &Password{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
