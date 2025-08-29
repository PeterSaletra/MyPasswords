package shell

import (
	"mypasswords/crypto"
	"mypasswords/store"

	"github.com/spf13/cobra"
)

type Shell struct {
	rootCmd *cobra.Command
	db      *store.Database
	keys    *crypto.Keys
}

func NewShell(db *store.Database, keys *crypto.Keys) *Shell {
	return &Shell{
		db:   db,
		keys: keys,
	}
}

func (s *Shell) PrepareCommands() {
	s.rootCmd = &cobra.Command{}

	s.rootCmd.AddCommand(&cobra.Command{
		Use:   "exit",
		Short: "Exit the Client",
		Long:  "Closes all clients and exits the CLI.",
		Run:   s.ExitCmd,
	})
	s.rootCmd.AddCommand(&cobra.Command{
		Use:   "clear",
		Short: "Clear the screen",
		Long:  "Clears the terminal screen.",
		Run:   s.ClearScreenCmd,
	})

	add := &cobra.Command{
		Use:   "add",
		Short: "Add a new password",
		Long:  "Adds a new password to the database.",
		Run:   s.AddPassword,
	}
	add.Flags().StringP("name", "n", "", "Name of the password entry")
	add.Flags().StringP("username", "u", "", "Username for the password entry")
	add.Flags().StringP("password", "p", "", "Password for the password entry")
	add.Flags().StringP("url", "l", "", "URL for the password entry")
	add.Flags().StringP("notes", "o", "", "Notes for the password entry")

	s.rootCmd.AddCommand(add)

}

func (s *Shell) Execute(args []string) error {
	s.rootCmd.SetArgs(args)
	return s.rootCmd.Execute()
}
