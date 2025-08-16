package shell

import (
	"mypasswords/store"

	"github.com/spf13/cobra"
)

type Shell struct {
	rootCmd *cobra.Command
	db      *store.Database
}

func NewShell(db *store.Database) *Shell {
	return &Shell{
		db: db,
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
}

func (s *Shell) Execute(args []string) error {
	s.rootCmd.SetArgs(args)
	return s.rootCmd.Execute()
}
