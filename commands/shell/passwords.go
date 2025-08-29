package shell

import (
	"mypasswords/crypto"
	"mypasswords/store"

	"github.com/spf13/cobra"

	"fmt"
)

func (s *Shell) AddPassword(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	username := cmd.Flag("username").Value.String()
	password := cmd.Flag("password").Value.String()
	url := cmd.Flag("url").Value.String()
	notes := cmd.Flag("notes").Value.String()

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
		Url:               url,
		Notes:             notes,
	}

	if err := s.db.CreatePassword(password_entry); err != nil {
		cmd.PrintErrf("Failed to add password: %v\n", err)
		return
	}

	cmd.Printf("Password added successfully: %s\n", name)
}
