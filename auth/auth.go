package auth

import (
	"fmt"
	"mypasswords/crypto"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/term"
)

type Keys struct {
	Master_hash string
	Master_key  string
}

func Authenticate(app string) (*Keys, error) {
	var keys Keys

	u, err := user.Current()
	if err != nil {
		return &keys, fmt.Errorf("Failed to get current user: %w", err)
	}
	username := u.Username

	// Add verification if this is first time using the app
	firstTime := checkUserFistTime()
	if firstTime {
		// Add hanlder for first time
		// Must include creating folder with subfolders for session and database
		// Propmting to enter new master password with double check
		password, err := handleFirstTime()
		if err != nil {
			return &keys, err
		}
		keys.Master_hash = crypto.DeriveMasterHash(password, username)
		keys.Master_key = crypto.DeriveMasterKey(keys.Master_hash, password)

		return &keys, nil

	} else {
		if ts, err := readSessionTimestamp(); err == nil {
			if time.Since(ts) < sessionDuration {
				return &keys, nil // Session still valid
			}
		}
	}

	fmt.Print("Enter new master password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))

	keys.Master_hash = crypto.DeriveMasterHash(string(password), username)
	keys.Master_key = crypto.DeriveMasterKey(keys.Master_hash, string(password))

	_ = saveSessionTimestamp(time.Now())
	return &keys, nil
}

func checkUserFistTime() bool {
	homedDir, err := os.UserHomeDir()
	if err != nil {
		return true
	}

	sessionFile := filepath.Join(homedDir, ".local", ".mypasswords", "session", "session_timestamp")
	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		return true
	}

	databasefile := filepath.Join(homedDir, ".local", ".mypasswords", "db", "mypasswords.db")
	if _, err := os.Stat(databasefile); os.IsNotExist(err) {
		return true
	}
	return false
}

func handleFirstTime() (string, error) {
	// u, err := user.Current()
	// if err != nil {
	// 	return err
	// }

	homedDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	sessionDir := filepath.Join(homedDir, ".local", ".mypasswords", "session")
	dbDir := filepath.Join(homedDir, ".local", ".mypasswords", "db")

	if err := os.MkdirAll(sessionDir, 0700); err != nil {
		return "", err
	}
	if err := os.MkdirAll(dbDir, 0700); err != nil {
		return "", err
	}

	// Prompt for master password
	fmt.Print("Enter new master password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	fmt.Print("Confirm new master password: ")
	confirmPassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	if string(password) != string(confirmPassword) {
		return "", fmt.Errorf("passwords do not match")
	}

	return "", saveSessionTimestamp(time.Now())
}
