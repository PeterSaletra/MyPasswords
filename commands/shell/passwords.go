package shell

import (
	"bufio"
	"mypasswords/crypto"
	"mypasswords/store"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"golang.org/x/term"

	"fmt"
)

func (s *Shell) AddPassword(cmd *cobra.Command, args []string) {

	reader := bufio.NewReader(os.Stdin)

	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		fmt.Print("Enter name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
	}

	username, _ := cmd.Flags().GetString("username")
	if username == "" {
		fmt.Print("Enter username: ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
	}

	password, _ := cmd.Flags().GetString("password")
	if password == "" {
		fmt.Print("Wanna generate? (Y/n): ")
		password, _ = reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if password == "Y" {
			// Add generating password
			password = "Super randomly generated password"
		} else {
			fmt.Print("Enter your password: ")
			passwordBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
			password = string(passwordBytes)
			password = strings.TrimSpace(password)
		}
	}

	url, _ := cmd.Flags().GetString("url")
	if url == "" {
		fmt.Print("Enter url: ")
		url, _ = reader.ReadString('\n')
		url = strings.TrimSpace(url)
	}

	notes, _ := cmd.Flags().GetString("notes")

	iv, err := crypto.GenerateIV(16)
	if err != nil {
		fmt.Print("Error occured while generating IV: %w", err)
	}
	encrypted, err := s.keys.EncryptAESCBC(password, iv)
	if err != nil {
		cmd.PrintErrf("Failed to encrypt password: %v\n", err)
		return
	}

	password_entry := &store.Password{
		Name:              name,
		Username:          username,
		EncryptedPassword: []byte(encrypted),
		IV:                iv,
		Url:               url,
		Notes:             notes,
		LastUsed:          time.Now(),
	}

	if err := s.db.CreatePassword(password_entry); err != nil {
		cmd.PrintErrf("Failed to add password: %v\n", err)
		return
	}

	cmd.Printf("Password added successfully: %s\n", name)
}

func (s *Shell) ListPasswords(cmd *cobra.Command, args []string) {
	passwords, err := s.db.GetAllPasswordsNames()
	if err != nil {
		cmd.PrintErrf("Failed to retrieve passwords: %v\n", err)
		return
	}

	for i, name := range passwords {
		cmd.Printf("  %d: %s\n", i+1, name)
	}
}

func (s *Shell) GetPassword(cmd *cobra.Command, args []string) {
	name := args[0]

	password, err := s.db.GetPasswordByName(name)
	if err != nil {
		cmd.Printf("Failed to retrieve password: %s\n", name)
		return
	}

	copy, _ := cmd.Flags().GetString("copy")
	cmd.Printf("Copying %s to clipboard\n", copy)
	if copy != "" {
		if err := clipboard.Init(); err != nil {
			cmd.PrintErrf("Failed to initialize clipboard: %v\n", err)
			return
		}
		switch copy {
		case "password":
			decrypted, err := s.keys.DecryptAESCBC(string(password.EncryptedPassword), password.IV)
			if err != nil {
				cmd.PrintErrf("Failed to decrypt password: %v\n", err)
				return
			}
			clipboard.Write(clipboard.FmtText, []byte(decrypted))

		case "username":
			clipboard.Write(clipboard.FmtText, []byte(password.Username))
		case "url":
			clipboard.Write(clipboard.FmtText, []byte(password.Url))
		default:
			cmd.PrintErrf("Invalid copy option: %s\n", copy)
			return
		}

		password.LastUsed = time.Now()
		if err = s.db.UpdatePassword(password); err != nil {
			cmd.PrintErrf("Failed to update password: %v\n", err)
			return
		}
	}

	show, _ := cmd.Flags().GetBool("show")
	if show {
		decrypted, err := s.keys.DecryptAESCBC(string(password.EncryptedPassword), password.IV)
		if err != nil {
			cmd.PrintErrf("Failed to decrypt password: %v\n", err)
			return
		}
		cmd.Printf("  Name: %s\n", password.Name)
		cmd.Printf("  Username: %s\n", password.Username)
		cmd.Printf("  Decrypted password: %s\n", decrypted)
		cmd.Printf("  Url: %s\n", password.Url)
		cmd.Printf("  Notes: %s\n", password.Notes)
		cmd.Printf("  Last Used: %s\n", password.LastUsed.Format(time.RFC1123))

		password.LastUsed = time.Now()
		if err = s.db.UpdatePassword(password); err != nil {
			cmd.PrintErrf("Failed to update password: %v\n", err)
			return
		}

		return
	}

	cmd.Printf("  Name: %s\n", password.Name)
	cmd.Printf("  Username: %s\n", password.Username)
	cmd.Printf("  Url: %s\n", password.Url)

	password.LastUsed = time.Now()
	if err = s.db.UpdatePassword(password); err != nil {
		cmd.PrintErrf("Failed to update password: %v\n", err)
		return
	}
}
