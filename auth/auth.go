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

type AuthResult struct {
	Keys      *crypto.Keys
	FirstTime bool
	Username  string
}

func Authenticate(app string) (*AuthResult, error) {
	var keys crypto.Keys

	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current user: %w", err)
	}
	username := u.Username

	firstTime := checkUserFistTime()
	if firstTime {
		password, err := handleFirstTime()
		if err != nil {
			return &AuthResult{Keys: &keys, FirstTime: true}, err
		}
		keys.DeriveMasterHash(password, username)
		keys.DeriveMasterKey(password)

		return &AuthResult{
			Keys:      &keys,
			FirstTime: true,
			Username:  username,
		}, nil
	}
	// } else {
	// 	if ts, err := readSessionTimestamp(); err == nil {
	// 		if time.Since(ts) < sessionDuration {
	// 			return &AuthResult{
	// 				Keys:      &keys,
	// 				FirstTime: false,
	// 				Username:  username,
	// 			}, nil
	// 		}
	// 	}
	// }

	fmt.Print("Enter new master password: ")
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))

	keys.DeriveMasterHash(string(password), username)
	keys.DeriveMasterKey(string(password))

	_ = saveSessionTimestamp(time.Now())
	return &AuthResult{
		Keys:      &keys,
		FirstTime: false,
		Username:  username,
	}, nil
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

	err = saveSessionTimestamp(time.Now())
	if err != nil {
		return "", err
	}

	return string(password), nil
}
