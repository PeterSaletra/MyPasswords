package cli

import (
	"github.com/chzyer/readline"
)

func GetConfig(menu string) *readline.Config {
	var config *readline.Config

	switch menu {
	case "main":
		config = &readline.Config{
			Prompt:          "-> ",
			HistoryFile:     "/tmp/readline.tmp",
			AutoComplete:    GetCompleter(menu),
			InterruptPrompt: "^C",
			EOFPrompt:       "exit",

			HistorySearchFold:   true,
			FuncFilterInputRune: filterInput,
		}

	}

	return config
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
