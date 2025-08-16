package commands

import (
	"fmt"
	"mypasswords/auth"
	"mypasswords/cli"
	"mypasswords/store"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mypasswords",
	Short: "MyPasswords CLI",
	Long:  `MyPasswords is a command line interface for managing your passwords.`,
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if err := auth.Authenticate("mypasswords"); err != nil {
		fmt.Fprintln(os.Stderr, "Authentication failed:", err)
		os.Exit(1)
	}

	db := store.NewDatabase()
	if err := db.Connect("your_master_key_here"); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to database:", err)
		return
	}

	c, err := cli.NewCli(db)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	c.Start()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
