package shell

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func (s *Shell) ExitCmd(cmd *cobra.Command, args []string) {

	fmt.Println("\nExiting QC2 Client. Goodbye!")
	time.Sleep(700 * time.Millisecond)

	fmt.Print("\033[H\033[2J")
	os.Exit(0)
}

func (s *Shell) ClearScreenCmd(cmd *cobra.Command, args []string) {
	fmt.Print("\033[H\033[2J")
}
