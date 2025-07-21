package cli

import (
	"fmt"
	"os/user"

	"github.com/chzyer/readline"
)

func GetConfig(menu string) (*readline.Config, error) {
	var config *readline.Config

	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}
	username := u.Username

	switch menu {
	case "main":
		config = &readline.Config{
			Prompt:          "\033[32m" + "@" + username + " -> " + "\033[0m",
			HistoryFile:     "/tmp/readline.tmp",
			AutoComplete:    GetCompleter(menu),
			InterruptPrompt: "^C",
			EOFPrompt:       "exit",

			HistorySearchFold:   true,
			FuncFilterInputRune: filterInput,
		}

	}

	return config, nil
}

func GetCompleter(menu string) *readline.PrefixCompleter {

	var completer *readline.PrefixCompleter

	return completer
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	case readline.CharCtrlL:
		if r == readline.CharCtrlL {
			print("\033[H\033[2J")
			return r, false
		}
	}
	return r, true
}
