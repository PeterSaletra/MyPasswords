package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/term"
)

func Authenticate(app string) error {
	// Add verification if this is first time using the app
	firstTime := checkUserFistTime()
	if firstTime {
		// Add hanlder for first time
		// Must include creating folder with subfolders for session and database
		// Propmting to enter new master password with double check
		return handleFirstTime()
	} else {
		if ts, err := readSessionTimestamp(); err == nil {
			if time.Since(ts) < sessionDuration {
				return nil // Session still valid
			}
		}
	}

	// u, err := user.Current()
	// if err != nil {
	// 	return fmt.Errorf("failed to get current user: %w", err)
	// }
	// username := u.Username

	_ = saveSessionTimestamp(time.Now())
	return nil
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

func handleFirstTime() error {
	// u, err := user.Current()
	// if err != nil {
	// 	return err
	// }

	homedDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	sessionDir := filepath.Join(homedDir, ".local", ".mypasswords", "session")
	dbDir := filepath.Join(homedDir, ".local", ".mypasswords", "db")

	if err := os.MkdirAll(sessionDir, 0700); err != nil {
		return err
	}
	if err := os.MkdirAll(dbDir, 0700); err != nil {
		return err
	}

	// Prompt for master password
	fmt.Print("Enter new master password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	fmt.Print("Confirm new master password: ")
	confirmPassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	if string(password) != string(confirmPassword) {
		return fmt.Errorf("passwords do not match")
	}

	return saveSessionTimestamp(time.Now())
}
