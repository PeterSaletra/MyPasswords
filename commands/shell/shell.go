package shell

import (
	"mypasswords/crypto"
	"mypasswords/store"

	"github.com/spf13/cobra"
)

type Shell struct {
	db   *store.Database
	keys *crypto.Keys
}

func NewShell(db *store.Database, keys *crypto.Keys) *Shell {
	return &Shell{
		db:   db,
		keys: keys,
	}
}

func (s *Shell) Execute(args []string) error {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "exit",
		Short: "Exit the Client",
		Long:  "Closes all clients and exits the CLI.",
		Run:   s.ExitCmd,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "clear",
		Short: "Clear the screen",
		Long:  "Clears the terminal screen.",
		Run:   s.ClearScreenCmd,
	})

	add := &cobra.Command{
		Use:   "add",
		Short: "Add a new password",
		Long:  "Adds a new password to the database interactively. Use flags to skip prompts.",
		Run:   s.AddPassword,
		Args:  cobra.NoArgs,
	}
	add.Flags().StringP("name", "n", "", "Name of the password entry")
	add.Flags().StringP("username", "u", "", "Username for the password entry")
	add.Flags().StringP("password", "p", "", "Password for the password entry")
	add.Flags().StringP("url", "l", "", "URL for the password entry")
	add.Flags().StringP("notes", "o", "", "Notes for the password entry")

	rootCmd.AddCommand(add)

	list := &cobra.Command{
		Use:   "list",
		Short: "List all passwords",
		Long:  "Lists all passwords stored in the database.",
		Run:   s.ListPasswords,
	}

	list.Flags().BoolP("show", "s", false, "Show detailed information")
	rootCmd.AddCommand(list)

	get := &cobra.Command{
		Use:   "get [name]",
		Short: "Get a password",
		Long:  "Retrieves a password from the database.",
		Args:  cobra.ExactArgs(1),
		Run:   s.GetPassword,
	}
	get.Flags().BoolP("show", "s", false, "Show detailed information")
	get.Flags().StringP("copy", "c", "", "Copy to clipboard: password, username, url")

	rootCmd.AddCommand(get)

	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
