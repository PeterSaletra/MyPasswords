package auth

import (
	"os"
	"path/filepath"
	"time"
)

func getSessionFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".local/.mypasswords/session", "session_timestamp"), nil
}

const sessionDuration = 15 * time.Minute

func readSessionTimestamp() (time.Time, error) {
	sessionFile, err := getSessionFile()
	if err != nil {
		return time.Time{}, err
	}
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return time.Time{}, err
	}
	ts, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return time.Time{}, err
	}
	return ts, nil
}

func saveSessionTimestamp(ts time.Time) error {
	sessionFile, err := getSessionFile()
	if err != nil {
		return err
	}
	return os.WriteFile(sessionFile, []byte(ts.Format(time.RFC3339)), 0600)
}
