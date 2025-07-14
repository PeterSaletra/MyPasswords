package commands

import (
	"fmt"
	"mypasswords/auth"
	"mypasswords/cli"
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
	c, err := cli.NewCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	c.Start()
}

func Execute() {
	if err := auth.Authenticate("mypasswords"); err != nil {
		fmt.Fprintln(os.Stderr, "Authentication failed:", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
