package auth

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/msteinert/pam"
	"golang.org/x/term"
)

func Authenticate(app string) error {
	// Check session file
	if ts, err := readSessionTimestamp(); err == nil {
		if time.Since(ts) < sessionDuration {
			return nil // Session still valid
		}
	}

	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}
	username := u.Username

	tx, err := pam.StartFunc(app, username, pamConv)
	if err != nil {
		return err
	}

	if err = tx.Authenticate(0); err != nil {
		fmt.Println("🔐 Cache wygasł, proszę podać hasło...")
		if err = tx.Authenticate(pam.Silent); err != nil {
			return err
		}
	}

	if err = tx.OpenSession(0); err != nil {
		return err
	}
	defer func() {
		if cerr := tx.CloseSession(0); cerr != nil {
			fmt.Fprintf(os.Stderr, "error closing PAM session: %v\n", cerr)
		}
	}()

	_ = saveSessionTimestamp(time.Now())
	return nil
}

func pamConv(style pam.Style, msg string) (string, error) {
	switch style {
	case pam.PromptEchoOff:
		fmt.Print(msg)
		buf, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		return string(buf), err
	default:
		fmt.Print(msg)
		var in string
		fmt.Scanln(&in)
		return in, nil
	}
}
